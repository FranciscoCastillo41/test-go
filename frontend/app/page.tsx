"use client";

import { useEffect, useState } from "react";

type HelloRes = { message: string; timestamp: string };
type EchoRes = { received: Record<string, unknown>; note: string };

export default function Home() {
    const [hello, setHello] = useState<HelloRes | null>(null);
    const [name, setName] = useState("Francisco");
    const [echo, setEcho] = useState<EchoRes | null>(null);
    const base = process.env.NEXT_PUBLIC_API_URL || ""; // empty -> dev proxy /api/*

    useEffect(() => {
        fetch(`${base}/api/hello`, { cache: "no-store" })
            .then((r) => r.json())
            .then(setHello)
            .catch((e) => console.error("hello error", e));
    }, [base]);

    async function submitEcho(e: React.FormEvent) {
        e.preventDefault();
        setEcho(null);
        const res = await fetch(`${base}/api/echo`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ name })
        });
        const data = (await res.json()) as EchoRes;
        setEcho(data);
    }

    return (
        <main style={{ padding: 24, fontFamily: "ui-sans-serif, system-ui" }}>
            <h1 style={{ fontSize: 28, marginBottom: 12 }}>Full-stack MVP (Next.js + Go)</h1>

            <section style={{ marginBottom: 24 }}>
                <h2 style={{ fontSize: 18 }}>Hello endpoint</h2>
                <pre style={{ background: "#111", color: "#0f0", padding: 12, borderRadius: 8 }}>
                    {hello ? JSON.stringify(hello, null, 2) : "Loading..."}
                </pre>
            </section>

            <section>
                <h2 style={{ fontSize: 18 }}>Echo test</h2>
                <form onSubmit={submitEcho} style={{ display: "flex", gap: 8, marginBottom: 12 }}>
                    <input
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                        placeholder="Enter a name"
                        style={{ padding: 8, border: "1px solid #ccc", borderRadius: 6 }}
                    />
                    <button type="submit" style={{ padding: "8px 14px", borderRadius: 6 }}>
                        Send
                    </button>
                </form>
                <pre style={{ background: "#111", color: "#0f0", padding: 12, borderRadius: 8 }}>
                    {echo ? JSON.stringify(echo, null, 2) : "Submit the form to see a response"}
                </pre>
            </section>

            <p style={{ marginTop: 24, opacity: 0.7 }}>
                API base: <code>{base || "dev proxy → http://localhost:8080/v1"}</code>
            </p>
        </main>
    );
}
