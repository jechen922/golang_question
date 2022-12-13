# 問題

1. service 繼承方法共用
   
   - 在 service pkg 中，讓各個 service 繼承 utils.ErrorHandler
   - 在回傳資料不存在時候，直接調用 handler 方法 NotFoundOrInternalError，
   方法如果收到 not found 以外的錯誤會直接判斷為內部錯誤，並且記錄起來
   - 某些商業邏輯在判斷是否資源存在時，則使用 ErrorOrInternalError 來處理允許資源存在或不存在

   設計是希望遇到相同情境可以直接使用共用方法，也可以減少寫測試案例。
   
   想問這樣的設計，有存在壞味道嗎？ 如果有的話，會建議怎麼調整呢？

2. apierror pkg 省略參數
   
   在 apierror pkg 方法 New，裡面讓 err 使用 ...error 方式傳入，但實際只用到第一筆 error

   如果不使用 `...error` 則需要像下方範例一樣傳入 nil
   ```go
   return apierror.New(apierror.InternalServiceError, nil)
   ```
      
   想問這樣以取巧的方式讓 function caller 可以省略參數，是個可接受的設計方式嗎？
   或是這種設計有埋藏其他的 bug 風險嗎？