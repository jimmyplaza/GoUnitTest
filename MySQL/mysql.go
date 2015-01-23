package main
import (
    "database/sql" //這包一定要引用，是底層的sql驅動
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "strconv" //這個是為了把int轉換為string
)

func main() { 
	db, err := sql.Open( "mysql" , "root:123456@tcp(localhost:3306)/test?charset=utf8" )
    if err != nil { 
    	panic(err.Error()) 
        fmt.Println(err.Error())
    }
    defer db.Close()  //只有在前面用了panic這時defer才能起作用，如果鏈接數據的時候出問題，他會往err寫數據
                                                                                 
    rows, err := db.Query( "select id,lvs from jimmy" )
    //判斷err是否有錯誤的數據，有err數據就顯示panic的數據
    if err != nil {
        panic(err.Error())
        fmt.Println(err.Error())
        return
    }
    defer rows.Close()
    var id int  
    var lvs string
    for rows.Next() { 
        rerr := rows.Scan(&id, &lvs)  //數據指針，會把得到的數據，往剛才id和lvs引入
        if rerr == nil {
            fmt.Println( "id號是" ,strconv.Itoa(id) + "lvs是" + lvs) //輸出來而已，看看
        }
    }
    insert_sql := "INSERT INTO jimmy(lvs) VALUES(?)"
    _, e4 := db.Exec(insert_sql, "nima")                                                            
    fmt.Println(e4)
    db.Close() //關閉數據庫
}