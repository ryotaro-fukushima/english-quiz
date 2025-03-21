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
    const [result, setResult] = useState<{ score: number; message: string } | null>(null);
    const router = useRouter();

    useEffect(() => {
        fetch("http://localhost:8080/api/questions") // GoのAPI
            .then((res) => res.json())
            .then((data) => setQuestions(data))
            .catch((error) => console.error("Error fetching questions:", error));
    }, []);

    const handleChange = (word: string, value: string) => {
        setAnswers((prev) => ({ ...prev, [word]: value }));
    };

    const handleSubmit = async () => {
        const response = await fetch("http://localhost:8080/api/check", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(answers),
        });

        const data = await response.json();
        setResult({ score: data.score, message: data.message });
    };

    return (
        <div className="min-h-screen flex flex-col items-center justify-center p-8">
            <h1 className="text-2xl font-bold mb-4">
                {mode === "en-to-jp" ? "英語 → 日本語" : "日本語 → 英語"}
            </h1>

            <ul className="space-y-4">
                {questions.map((q, index) => (
                    <li key={index} className="border p-4 rounded-md flex flex-col">
                        <strong className="mb-2">
                            {mode === "en-to-jp" ? q.word : q.translations[0]}
                        </strong>
                        <input
                            type="text"
                            value={answers[q.word] || ""}
                            onChange={(e) => handleChange(q.word, e.target.value)}
                            className="border p-2 rounded-md"
                            placeholder="答えを入力"
                        />
                    </li>
                ))}
            </ul>

            <button
                onClick={handleSubmit}
                className="mt-4 bg-blue-500 text-white px-4 py-2 rounded-md"
            >
                送信
            </button>

            {result && (
                <div className="mt-4 p-4 border rounded-md">
                    <p>スコア: {result.score}点</p>
                    <p>{result.message}</p>
                </div>
            )}

            <button
                onClick={() => router.push("/")}
                className="mt-4 bg-gray-500 text-white px-4 py-2 rounded-md"
            >
                ホームに戻る
            </button>
        </div>
    );
}
