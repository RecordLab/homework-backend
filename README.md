# **데일리 스쿱 (Daily Scoop)**
> 일기 작성 애플리케이션

하루 한 장 사진과 함께 쓰는 모바일 일기장

![재미(2)](https://user-images.githubusercontent.com/68420164/138399327-8ee88ef6-2c36-42d0-8be5-6567dd0f4855.png)

## 🐶 주요기능
+ 일기 작성
+ 테마 선택
+ 모아보기
+ 일기 통계
+ 일기 검색
+ 잠금 설정

## 🐱 세부기능

|구분|기능|설명|비고|
|------|---|---|---|
|1|일기 작성|갤러리에서 사진을 선택해서 추가하고 텍스트 에디터로 글을 쓸 수 있다||
|2|테마 선택|기분에 따른 감정 이모티콘과 일기 테마를 선택할 수 있다||
|3|모아보기|작성한 일기를 날짜별로 확인할 수 있다||
|4|일기 통계|일기 작성 시 선택한 감정들을 일자별로 모아 통계를 볼 수 있다|
|5|일기 검색|키워드를 통해 특정 일기를 검색할 수 있다|
|6|잠금 설정|앱 자체 잠금 설정을 통해 프라이버시를 지킬 수 있다|

## 🐭 아키텍처
![2021-10-22 14;51;47](https://user-images.githubusercontent.com/68420164/142414014-4d70294c-5146-4731-af14-074f87359bb7.png)

## 🐹 설치
```
플레이스토어 배포 예정
```

## 🐰 사용 예시
#### 로그인 & 회원가입
![image](https://user-images.githubusercontent.com/68420164/142501767-87f46fc6-8dc7-4533-8227-8ce2e01290be.png)
#### 메인 페이지 & 일기작성
![image](https://user-images.githubusercontent.com/68420164/142417358-d3c4276a-f020-4f13-b253-5b649a39fec1.png)
![image](https://user-images.githubusercontent.com/68420164/142417467-6ea15856-6747-4bc2-81e9-73b25b863e52.png)
#### 모아보기 & 검색 & 일기 상세페이지
![image](https://user-images.githubusercontent.com/68420164/142417598-be710529-13c8-40be-aa09-027bd7e7ed4e.png)
![image](https://user-images.githubusercontent.com/68420164/142417677-db870027-9d7e-4b8d-928d-69b017c1244b.png)
![image](https://user-images.githubusercontent.com/68420164/142417736-f3eea291-c171-48ef-ab59-bff915e66944.png)
#### 통계 & 프로필 페이지
![image](https://user-images.githubusercontent.com/68420164/142417923-3e7bac7a-fc7f-47d6-b0cb-cf0791ed278e.png)
![image](https://user-images.githubusercontent.com/68420164/142417968-e5b32208-ff4d-475b-a22e-3336d96a6f0b.png)
## 💻 버전 정보
#### Android
```
Android Studio Arctic Fox
Kotlin "1.5.31"
androidX Core "1.6.0"
compileSdkVersion 31
buildToolsVersion '30.0.3'
jvmTarget = '1.8'
```
#### Backend
```
Go 1.17.2
GoLand 2021.2.4
echo framework 4.6.1
MongoDB 5.0.3
Docker 20.10.10
Docker Comopose 1.29.2
```
## 🐵 기여
1. 해당 프로젝트를 Fork 하세요
2. feature 브랜치를 생성하세요 (git checkout -b feature/fooBar)
3. 변경사항을 commit 하세요 (git commit -am 'Add some fooBar')
4. 브랜치에 Push 하세요 (git push origin feature/fooBar)
5. 새로운 Merge Request를 요청하세요

## 🐧 개발 일정
## 1주(10.11 - 10.15)

### ✅ 기획

🔸 기획 의도 및 아이템 선정

🔸 필요 서비스 아이템 기획 회의

🔸 유사 서비스 조사

🔸 요구 명세서 제작

🔸 개발 환경 구축

## 2주(10.18 - 10.22)

### ✅ 기획

🔸 주요 기능 구상

🔸 부가 기능 구상 및 우선순위 설계

🔸 와이어 프레임 제작

🔸 원활한 소통 위한 슬랙, 노션 등의 채널 사용

## 3주(10.25 - 10.29) 

### ✅ ANDROID

🔸 Android 뷰 제작
```
🔸 [메인] 메인 화면 레이아웃 제작

🔸 [메인] 프로필 화면 레이아웃 제작

🔸 [기본] 다이얼로그 뷰 제작
```

### ✅ IOS

🔸 IOS 뷰 제작

```
🔸 [메인] 메인 페이지 레이아웃 제작

🔸 [일기] 일기 작성 페이지 레이아웃 제작
```

### ✅ Back

🔸 [DB] User DB 설계 및 구축

🔸 [DB] 일기 DB 설계 및 구축


## 4주(11.01 - 11.05) 

### ✅ ANDROID

🔸 Android 뷰 제작
```
🔸 [메인] 검색 화면 레이아웃 제작

🔸 [메인] 메인 페이지 더 보기 레이아웃 제작

🔸 [일기] 모아보기 화면 레이아웃 제작

🔸 [일기] 일기 상세 화면 레이아웃 제작

🔸 [일기] 일기 작성 화면 레이아웃 제작

🔸 [기본] 회원가입 화면 레이아웃 제작

🔸 [기본] 스플래시 화면 레이아웃 제작

🔸 [기본] 잠금 화면 레이아웃 제작

```

### ✅ Back

🔸 [유저] 회원 CRUD API 구현
```
🔸 [기본] 회원 가입 기능

🔸 [기본] 회원 정보 수정 기능

🔸 [기본] 회원 탈퇴 기능

🔸 [기본] 회원 정보 조회 기능

```
## 5주(11.08 - 11.12)

### ✅ ANDROID

🔸 Android Activity 제작
```
🔸 [메인] 홈 화면 캘린더 및 미리보기 구현

🔸 [일기] 서버와 연동해서 전체 일기 조회

🔸 [일기] 리사이클러뷰를 활용한 모아보기 화면 구현

🔸 [일기] 일기 상세보기 화면 구현

🔸 [일기] 일기 작성 화면 구현 및 통신 로직 구현

🔸 [기본] 로그인 및 회원가입 로직 구현

🔸 [기본] 스플래시 화면 구현 및 화면 분기 로직 구현

🔸 [기본] 잠금 화면 구현 및 암호 저장

🔸 [메인] 프로필 화면 구현 및 설정값 저장
```

### ✅ Back

🔸 [일기] 일기 CRUD API 구현
```
🔸 [기본] 일기 작성 기능

🔸 [기본] 일기 수정 기능

🔸 [기본] 일기 조회 기능

🔸 [기본] 일기 삭제 기능
```
🔸 [일기] 이미지 호스팅 구현

## 6주(11.15 - 11.19)

### ✅ ANDROID

🔸 Android Activity 제작
```
🔸 [메인] 일기 검색 기능 구현

🔸 [메인] 통계 화면 구현 및 통계 데이터 수신

🔸 [일기] 일기 수정 및 삭제 로직 구현

🔸 [기본] 로그아웃 및 회원탈퇴 로직 구현

🔸 [기본] UI 수정

```
### ✅ Back

```
🔸 [기본] 통계 데이터 API 구현

🔸 [기본] 유저 정보 DB 수정

🔸 [기본] 소셜 로그인

🔸 [기본] 일기 페이지네이션

🔸 [기본] 이미지 리사이징

```
