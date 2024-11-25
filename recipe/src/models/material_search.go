package models

import (
	"errors"
)

//カテゴリーからレシピを検索する関数
func Material_Search(uid string)(Material,error) {
    // 材料を定義
	material := &Material{
		Uid:uid,
	}
    
    //からidを取得
	material_len := dbconn.Where(material).Find(&material).RowsAffected

    //材料が見つからなかった場合
	if material_len == 0 { 
	    return *material,errors.New("Material not found")
	}

    return *material,nil
}

