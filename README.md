# 📚 英単語学習アプリ

英語 ⇔ 日本語のクイズ形式で英単語を学習できるアプリです！
**Next.js（TypeScript）+ Go（Gin）** を使用して構築。

## ✨ 特徴
- **英語 → 日本語** & **日本語 → 英語** の2つのモード
- 10問ランダム出題
- **大文字・小文字、スペースを無視** して正誤判定
- スコアによって「褒める or 励ます」メッセージ
- **Go + Gin で API** を構築
- **Next.js + TypeScript でフロントエンド**

---

## 🛠 使用技術
- **フロントエンド:** Next.js (TypeScript, TailwindCSS)
- **バックエンド:** Go (Gin)
- **データ:** ハードコードされた英単語リスト（今後データベース対応予定）

---

## 🚀 インストール & 実行

### 1️⃣ **バックエンド（Go）**
```sh
cd englishapp
go run main.go
```

### 2️⃣ **フロントエンド（Next.js）**
```sh
cd english-quiz-frontend
npm install
npm run dev
