cách sử dụng
-tạo database source tên "okr" trong postgres với mật khẩu "den"
(có thể tùy chỉnh database source trong file config.env ) 
-trong postgres:
    + \c ork (sử dụng database ork)
    + create extension if not exist "uuid-ossp";(cài đặt uuid-ossp )
-trong vscode:
    + go run . init 
    + go run . import -dir ITDD (import toàn bộ file exel trong folder ITD vào Postgres)
    
**Nếu sheet exel của user không đủ thông tin thì sheet sẽ bị bỏ 
**Sẽ xuất hiện rất nhiều manager_ID có trùng tên trong bảng OKR_Users
!^! Đề xuất :
- cho người dùng mẫu exel 


