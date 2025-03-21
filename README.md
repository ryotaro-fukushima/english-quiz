# 📚 英単語学習アプリ

英語 ⇔ 日本語のクイズ形式で英単語を学習できるアプリです！  
**Next.js（TypeScript）+ Go（Gin）** を使用して構築。

---

## ✨ 特徴

✅ **英語 → 日本語** & **日本語 → 英語** の 2 つのモード  
✅ **10 問ランダム出題**（毎回違う問題が出る！単語は随時追加予定）  
✅ **大文字・小文字、スペースを無視** して正誤判定  
✅ **スコアによって「褒める or 励ます」メッセージ**  
✅ **スコア履歴をローカル保存！過去のスコアが見られる**  
✅ **Go + Gin で API を構築** / **Next.js + TypeScript でフロントエンド**

---

## 🛠 使用技術

| **技術**                 | **使用目的**        |
| ------------------------ | ------------------- |
| **Next.js** (TypeScript) | フロントエンド      |
| **TailwindCSS**          | スタイリング        |
| **Go (Gin)**             | バックエンド（API） |
| **localStorage**         | スコア履歴保存      |

---

## 🚀 **インストール & 実行方法**

### **1️⃣ バックエンド（Go）**

```sh
cd english-quiz
go run main.go
```

### 2️⃣ **フロントエンド（Next.js）**

```sh
cd english-quiz-frontend
npm install
npm run dev
```
