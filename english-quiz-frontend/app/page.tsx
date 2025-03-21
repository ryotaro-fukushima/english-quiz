"use client";

import { useRouter } from "next/navigation";

export default function Home() {
  const router = useRouter();

  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-gray-100 p-8">
      {/* タイトル */}
      <h1 className="text-4xl font-extrabold text-gray-900 mb-4">📚 英単語クイズ</h1>
      <p className="text-lg text-gray-700 mb-8">出題モードを選んでください</p>

      {/* ボタンコンテナ */}
      <div className="flex flex-col space-y-4">
        {/* 英語 → 日本語 */}
        <button
          onClick={() => router.push("/en-to-jp")}
          className="w-64 h-16 bg-blue-500 text-white rounded-lg text-xl font-semibold shadow-lg hover:bg-blue-600 transition"
        >
          英語 → 日本語
        </button>

        {/* 日本語 → 英語 */}
        <button
          onClick={() => router.push("/jp-to-en")}
          className="w-64 h-16 bg-green-500 text-white rounded-lg text-xl font-semibold shadow-lg hover:bg-green-600 transition"
        >
          日本語 → 英語
        </button>
      </div>
    </div>
  );
}
