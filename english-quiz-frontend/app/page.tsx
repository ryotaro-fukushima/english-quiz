"use client";

import { useRouter } from "next/navigation";

export default function Home() {
  const router = useRouter();

  return (
    <div className="min-h-screen flex flex-col items-center justify-center p-8">
      <h1 className="text-2xl font-bold mb-6">英単語クイズ</h1>
      <p className="mb-4">出題モードを選んでください</p>

      {/* 英語 → 日本語ボタン */}
      <button
        onClick={() => router.push("/en-to-jp")}
        className="bg-blue-500 text-white px-6 py-3 rounded-md text-lg mb-4"
      >
        英語 → 日本語
      </button>

      {/* 日本語 → 英語ボタン */}
      <button
        onClick={() => router.push("/jp-to-en")}
        className="bg-green-500 text-white px-6 py-3 rounded-md text-lg"
      >
        日本語 → 英語
      </button>
    </div>
  );
}
