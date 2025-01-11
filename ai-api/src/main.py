from bs4 import BeautifulSoup
import json
import requests
import uvicorn
from fastapi import FastAPI,Request
from fastapi.responses import JSONResponse
from typing import List, Optional
from pydantic import BaseModel
import os
import google.generativeai as genai
from dotenv import load_dotenv

# .envファイルの内容を読み込見込む
load_dotenv()

from slowapi.errors import RateLimitExceeded
from slowapi import Limiter, _rate_limit_exceeded_handler
from slowapi.util import get_remote_address

# APIキー設定
apiKey = os.getenv("GENAI_API_KEY")
genai.configure(api_key=apiKey)

# サーバー作成
limiter = Limiter(key_func=get_remote_address)
app = FastAPI()
app.state.limiter = limiter
app.add_exception_handler(RateLimitExceeded, _rate_limit_exceeded_handler)

class PostURL(BaseModel):
    Url: str

@app.post("/recipe")
@limiter.limit("1/second")
async def CovertRecipe(request: Request,args: PostURL):
    try:
        # リクエスト送信
        reqUrl = args.Url

        htmlContent = requests.get(reqUrl).content

        soup = BeautifulSoup(htmlContent.decode('utf-8'), 'html.parser')

        jsons = soup.find_all("script", {"type": "application/ld+json"})

        for extractjson in jsons:
            result = json.loads(extractjson.text)

            # @type で判定
            if result["@type"] == "Recipe":
                print(result)
                break

        # Create the model
        generation_config = {
            "temperature": 0.8,
            "top_p": 0.95,
            "top_k": 40,
            "max_output_tokens": 8192,
            "response_mime_type": "application/json",
        }

        model = genai.GenerativeModel(
            model_name="gemini-1.5-flash-8b",
            generation_config=generation_config,
        )

        chat_session = model.start_chat(
            history=[
                {
                    "role": "user",
                    "parts": [
                        "```json\n{\n    \"recipeCategory\": \"主菜\", //料理の種類が入ります [主菜,主食,副菜,汁物] のどれかです\n    \"recipeName\": \"野菜キーマカレー\", //料理の名前が入ります\n    \"recipeImage\": \"\", //画像は空白のままで大丈夫です\n    \"finalState\": \"Reheat\", //料理の最終状態 [Hot,Reheat,Cool,Normal] のどれかは入ります\n    \"steps\": [ //各料理の手順が入ります 配列です\n        {\n            \"name\": \"玉ねぎを見つめる\", //手順の名前です 手順の場合は(手順1,手順2) 下準備の場合は(下準備1,下準備2) 仕上げの場合は(仕上げ1,仕上げ2) などの連番で入れてください\n            \"time\": 5, //手順の所要時間です 大体の時間を入れてください 単位は分です\n            \"type\": \"下準備\", //作業の種類です [下準備,調理,仕上げ] のどれかが入ります\n            \"concurrent\": \"不可\", //他の作業と並行できるか出来ないかです [可,不可] のどちらかが入ります\n            \"ingredients\": [ //手順で使う材料が入ります 配列です\n                {\n                    \"name\": \"玉ねぎ\", //材料名です\n                    \"quantity\": 1, //材料の数です     (少量は適量の場合はおおよそのgに変換してください 1/2 などの場合も全ておおよそのgに変換してください)\n                    \"unit\": \"個\" //材料の単位です   (少量は適量の場合はおおよそのgに変換してください 1/2 などの場合も全ておおよそのgに変換してください)\n                }\n            ],\n            \"utensils\": [ //手順で使う道具が入ります 配列です\n                {\n                    \"name\": \"包丁\", //道具名です\n                    \"quantity\": 1, //道具の数です\n                    \"unit\": \"本\" //道具の単位です\n                }\n            ],\n            \"description\": \"玉ねぎをじっと見つめましょう\" //作業の説明です\n        }\n    ]\n}\n```\n\nYou need to process the json of the recipe\nPlease process the json passed to you.\nThe rules are as follows\nPlease refer to the above json for the format of the output.\nWhen creating the json, please follow the comments in the above json\nOnly output json\nDo not include json comments\nIf time is null, enter approximate time\nFor null, please enter “”.\nquantity must be a number only\nIn the ingredients name field, only the name of the ingredient should be extracted",
                    ],
                },
            ]
        )

        response = chat_session.send_message(json.dumps(result))
        result = json.loads(response.text)
        # recipeImage を削除
        del result["recipeImage"]

        return JSONResponse(content=result)
    except:
        import traceback
        traceback.print_exc()
    
        return ""

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000, log_level="debug")