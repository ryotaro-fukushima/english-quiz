"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";

type Question = {
    word: string;
    translations: string[];
};

export default function Quiz({ mode }: { mode: "en-to-jp" | "jp-to-en" }) {
    const [questions, setQuestions] = useState<Question[]>([]);
    const [answers, setAnswers] = useState<{ [key: string]: string }>({});
    const [result, setResult] = useState<{ score: number; message: string; results: { [key: string]: boolean } } | null>(null);
    const [scoreHistory, setScoreHistory] = useState<number[]>([]); // 🏆 スコア履歴
    const router = useRouter();

    useEffect(() => {
        fetch("http://localhost:8080/api/questions")
            .then((res) => res.json())
            .then((data) => setQuestions(data))
            .catch((error) => console.error("Error fetching questions:", error));

        // 🏆 スコア履歴を取得（初回読み込み時）
        const savedScores = JSON.parse(localStorage.getItem("scoreHistory") || "[]");
        setScoreHistory(savedScores);
    }, []);

    const handleChange = (key: string, value: string) => {
        setAnswers((prev) => ({ ...prev, [key]: value }));
    };

    const handleSubmit = async () => {
        console.log("送信するデータ:", answers);

        const formattedAnswers = Object.fromEntries(
            Object.entries(answers).map(([word, userInput]) => [
                word,
                userInput.trim().toLowerCase(),
            ])
        );

        const response = await fetch("http://localhost:8080/api/check", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(formattedAnswers),
        });

        const data = await response.json();
        console.log("受け取ったデータ:", data);

        setResult(data);

        // 🏆 履歴を更新
        const newHistory = [data.score, ...scoreHistory].slice(0, 3); // 最新3回分を保持
        setScoreHistory(newHistory);
        localStorage.setItem("scoreHistory", JSON.stringify(newHistory)); // ブラウザに保存
    };

    return (
        <div className="min-h-screen flex flex-col items-center justify-center bg-gray-100 p-6">
            <h1 className="text-3xl font-bold text-gray-900 mb-6">
                {mode === "en-to-jp" ? "英語 → 日本語" : "日本語 → 英語"}
            </h1>

            <div className="w-full max-w-lg bg-white p-6 rounded-lg shadow-md">
                <ul className="space-y-4">
                    {questions.map((q, index) => {
                        const key = mode === "en-to-jp" ? q.word : q.translations[0];

                        return (
                            <li key={index} className="border border-gray-300 p-4 rounded-lg flex flex-col bg-white shadow-md">
                                <strong className="mb-2 text-lg text-gray-900">{key}</strong>
                                <input
                                    type="text"
                                    value={answers[key] || ""}
                                    onChange={(e) => handleChange(key, e.target.value)}
                                    className="border border-gray-400 p-3 rounded-lg w-full text-lg text-gray-900 bg-white"
                                    placeholder="答えを入力"
                                />
                            </li>
                        );
                    })}
                </ul>
            </div>

            <button
                onClick={handleSubmit}
                className="mt-6 bg-blue-500 text-white px-6 py-3 rounded-lg text-lg shadow-md hover:bg-blue-600 transition"
            >
                解答する
            </button>

            {result && (
                <div className="mt-6 p-6 w-full max-w-lg bg-white rounded-lg shadow-lg text-center">
                    <p className="text-4xl font-bold text-gray-900 bg-blue-100 px-4 py-2 rounded-lg inline-block">
                        スコア: {result.score}点
                    </p>
                    <p className="mt-2 text-lg text-gray-700 font-medium">{result.message}</p>

                    {/* 🏆 スコア履歴を表示 */}
                    <div className="mt-6 p-4 bg-gray-100 rounded-lg shadow-md">
                        <h3 className="text-lg font-bold text-gray-800 mb-3">📊 過去のスコア</h3>
                        <ul className="space-y-2">
                            {scoreHistory.map((score, index) => (
                                <li
                                    key={index}
                                    className={`p-3 rounded-md text-lg font-medium text-center shadow-sm ${score === Math.max(...scoreHistory) ? "bg-green-100 text-green-700 font-bold" : "bg-gray-200 text-gray-900"
                                        }`}
                                >
                                    {score}点
                                </li>
                            ))}
                        </ul>
                    </div>

                    {/* 正誤リスト */}
                    <ul className="mt-4 space-y-2">
                        {questions.map((q, index) => {
                            const key = mode === "en-to-jp" ? q.word : q.translations[0];
                            const isCorrect = result.results[key];

                            return (
                                <li key={index} className={`flex justify-between items-center p-3 rounded-md shadow-sm ${isCorrect ? "bg-green-100" : "bg-red-100"}`}>
                                    <div className="flex items-center space-x-2">
                                        <span className="text-lg font-semibold text-gray-900">{key}</span>
                                        {!isCorrect && (
                                            <span className="text-sm text-blue-500">
                                                （正解: {mode === "en-to-jp" ? q.translations.join(", ") : q.word}）
                                            </span>
                                        )}
                                    </div>
                                    <span className={`text-2xl font-bold ${isCorrect ? "text-green-500 animate-bounce" : "text-red-500"}`}>
                                        {isCorrect ? "✔" : "✖"}
                                    </span>
                                </li>
                            );
                        })}
                    </ul>
                </div>
            )}

            <button
                onClick={() => router.push("/")}
                className="mt-4 bg-gray-500 text-white px-4 py-2 rounded-lg text-lg shadow-md hover:bg-gray-600 transition"
            >
                ホームに戻る
            </button>
        </div>
    );
}
