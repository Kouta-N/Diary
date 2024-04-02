//Validの真偽条件で=を抜かして、v = 0のときの場合漏れを起こした。

Id: sql.NullInt64{
	Valid: v.Id >= 0, // >0にしてバグった
	Int64: v.Id,
},