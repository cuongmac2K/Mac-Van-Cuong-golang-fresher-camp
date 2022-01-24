create database QLBH;

CREATE TABLE QLBH.Users (
  id_user INT NOT NULL,
  HoTen VARCHAR(45) ,
  NgaySinh datetime ,
  DiaChi NVARCHAR(50) ,
  NgayDangKi datetime ,
  PRIMARY KEY (id_user));
  
  
  create table QLBH.TaiKHoan (
  id_taiKhoan INT NOT NULL,
   id_user INT NOT NULL,
   TenDangNhap VARCHAR(50) not null,
   Pass VARCHAR(20) not null,
   PRIMARY KEY (id_taiKhoan),
   CONSTRAINT fk_user FOREIGN KEY (id_user) REFERENCES Users(id_user)
  );
  
  CREATE TABLE QLBH.PhongBan (
  id_PhongBan INT NOT NULL,
  TenPhongBan NVARCHAR(45) ,
  TruongPhong nvarchar(50) ,
  PRIMARY KEY (id_PhongBan));
  
    create table QLBH.NhanVien (
  id_NhanVien INT NOT NULL,
   id_PhongBan INT NOT NULL,
   HoTen VARCHAR(50) not null,
    NgaySinh datetime ,
  DiaChi NVARCHAR(50) ,
   PRIMARY KEY (id_NhanVien),
   CONSTRAINT fk_PhongBan FOREIGN KEY (id_PhongBan) REFERENCES PhongBan(id_PhongBan)
  );
  
   CREATE TABLE QLBH.CuaHang (
  id_CuaHang INT NOT NULL,
  DiaCHi NVARCHAR(45) ,
  TenCuaHang nvarchar(50) ,
  SoLuongNhanVien INT,
  NgQuanLy nvarchar(50),
  NgayBatDauHoatDong datetime,
  PRIMARY KEY (id_CuaHang));
   
  CREATE TABLE QLBH.CacPhuongThucThanhToan (
  id_LoaiPhuongThucThanhToan INT NOT NULL,
  TenPhuongThuc NVARCHAR(45) ,
  PRIMARY KEY (id_LoaiPhuongThucThanhToan));
  
     CREATE TABLE QLBH.ThanhToanQuaThe (
	id INT NOT NULL,
	TenNganHang NVARCHAR(45) ,
	id_LoaiPhuongThucThanhToan INT ,
	STK INT,
	TenChuTheNhan nvarchar(50),
    TenChuTheChuyen nvarchar(50),
    NgayMoTheCuaNgGui datetime,
    NoiDung nvarchar(200),
    ThoiGianThanhToan datetime,
    PRIMARY KEY (id),
    Constraint fk_CacPPTT Foreign key(id_LoaiPhuongThucThanhToan) References CacPhuongThucThanhToan(id_LoaiPhuongThucThanhToan)
	);
    
    Create table QLBH.NhaCungCap(
    id_NCC INT not null,
    TenNhaCC nvarchar(100),
    DiaChi nvarchar(100),
    NguoiDaiDien nvarchar(100),
    SDT int ,
    primary key(id_NCC));
    
  Create table QLBH.HangCungCap(
	id_HCC int not null,
	 id_NCC INT not null,
     TenHangCC varchar(100),
     SoLuong int,
     Dongia int,
	primary key(id_HCC),
    constraint fk_NCC foreign key(id_NCC) references NhaCungCap(id_NCC)
  );
  Create table QLBH.LoaiHang(
	id_loaiHang int not null,
    TenLoaiHang nvarchar(100),
    primary key(id_loaiHang)
  );
  create table QLBH.GiamGia(
  id_GiamGia int not null,
	TenMaGiamGia nvarchar(100),
    PhamTramGiam int,
    primary key(id_GiamGia)
  );
  create table QLBH.SanPham(
	id_sanPham int not null,
    id_HCC int not null,
    id_loaiHang int not null,
    TenHang nvarchar(100),
    SoLuongCo int,
    DonGia int,
    DonViDem nvarchar(100),
    MoTa nvarchar(200),
    primary key(id_sanPham) ,
    constraint fk_HCC foreign key (id_HCC) references HangCungCap(id_HCC),
    constraint fk_LoaiHang foreign key (id_loaiHang) references LoaiHang(id_loaiHang)
  );
  create table QLBH.Giohang(
  id_gioHang int not null,
  id_user int not null,
    id_GiamGia int not null,
    id_sanPham int not null,
    SoLuongSanPhamDaChon int,
    ThoiGianChon datetime,
    primary key(id_gioHang),
    constraint fk_User1 foreign key (id_user) references Users(id_user),
    constraint fk_giamGia foreign key (id_GiamGia) references GiamGia(id_GiamGia),
    constraint fk_SP foreign key (id_sanPham) references SanPham(id_sanPham)
  );
  create table QLBH.HoaDon(
  id_hoaDon int not null,
  	id_LoaiPhuongThucThanhToan INT ,
  listMaSp varchar(1000),
  ThanhTien Float ,
  primary key(id_hoaDon),
  constraint fk_PhuongThucThanhToan2 foreign key (id_LoaiPhuongThucThanhToan) references CacPhuongThucThanhToan(id_LoaiPhuongThucThanhToan)
  );
   
  