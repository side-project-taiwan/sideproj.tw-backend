/*
project api 欄位
1. logo圖
2. *專案 title
3. tags ( 逗點分隔 )
4. *專案敘述
5. *id
6. 是否釘選(pinned)
*/

CREATE TABLE project (
  pid 								INT AUTO_INCREMENT PRIMARY KEY,			# ppid (遞增編號)
  id 								CHAR(36) NOT NULL,  					# UUID (API使用)
  project_name 				        VARCHAR(255) NOT NULL,					# 專案名稱
  pinned							BOOL	NOT NULL,						# 釘選 (專案置頂)
  project_description               VARCHAR(255) NOT NULL,					# 專案敘述
  tags 								VARCHAR(255),							# 標籤 (逗點分隔 例如 pm, uiux, backend, frontend)
  logo_picture 				        VARCHAR(255),							# Logo 圖 (放在 cloudflare R2, default 一個 mock 圖片)
  github_url 					    VARCHAR(255),							# Github 連結
  site_url 						    VARCHAR(255),							# 網站連結
  owner_email 				        VARCHAR(255),							# 創辦人 Email
  update_at 					    DATETIME DEFAULT CURRENT_TIMESTAMP,		# 更改時間
  created_at 					    DATETIME DEFAULT CURRENT_TIMESTAMP		#	創建時間
);



/*
event api 欄位
1. logo圖
2. *專案 title
3. tags ( 逗點分隔 )
4. *專案敘述
5. *id
6. *discord活動連結
7. *活動時間
*/

CREATE TABLE event (
  eid 								INT AUTO_INCREMENT PRIMARY KEY,			# eid (遞增編號)
  id 								CHAR(36) NOT NULL,  					# UUID (API使用)
  event_name	 				    VARCHAR(255) NOT NULL,					# 活動名稱
  event_description 	            VARCHAR(255) NOT NULL,					# 活動敘述
  tags 								VARCHAR(255),							# 標籤 (逗點分隔 例如 pm, uiux, backend, frontend)
  picture 						    VARCHAR(255),							# 圖片 (放在 cloudflare R2, default 一個 mock 圖片)
  event_url 					    VARCHAR(255) NOT NULL,					# event 連結(目前都是 discord 頻道連結)
  event_speaker 			        VARCHAR(255) NOT NULL,					# 活動講者
  event_time					    DATETIME NOT NULL,						# 活動時間
  update_at 					    DATETIME DEFAULT CURRENT_TIMESTAMP,		# 更改時間
  created_at 					    DATETIME DEFAULT CURRENT_TIMESTAMP		#	創建時間
);