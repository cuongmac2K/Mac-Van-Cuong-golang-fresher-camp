Vì sao không nên chứa file upload vào ngay chính bên trong service mà nên dùng Cloud: 
  +lưu file upload ở trên cloud k giới hạn dung lượng,còn lưu trên service thì có giới hạn dung lượng.
  +khi người dùng sử dụng web thì có một số file người dùng lúc mới chạy web hoặc trong quá trình sử dụng sẽ k dùng tới hoặc dùng sau. nếu cho file vào trong service thì sẽ làm tăng dung lượng của service dẫn đến giảm tốc độ duyệt web.
  +lưu trên cloud thì ta có thể dùng lại file đó ở những dự án khác.
  +khi lưu trên cloud  thì có thể khử các file trùng lặp.

Vì sao không chứa binary ảnh vào DB:
  +khi lưu file binary vào DB sẽ làm DB nặng thêm đáng kể.
  +vì kiểu dữ liệu là dạng binary.dữ liệu dc tổ chức thành một hàng dài nên nó sẽ khó toàn vẹn khi xảy ra sự cố khi truyền file.
  +và khi lưu vào đó thì sau mỗi lần cập nhật thì dữ liệu cũ sẽ bị mất đi.nếu ta muốn dùng lại dự liệu đó thì k còn nữa.
  
