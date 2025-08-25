// app/error.tsx
"use client";

import { useEffect } from "react";

export default function Error({
  error,
  reset,
}: {
  error: Error & { digest?: string };
  reset: () => void;
}) {
  useEffect(() => {
    // Report to your logger/SaaS here
    console.error("Route error:", error);
  }, [error]);

  return (
    <main>
      <h1>Something went wrong</h1>
      <p style={{ color: "crimson" }}>{error.message}</p>
      <button onClick={() => reset()} style={{ marginTop: 12 }}>
        Try again
      </button>
    </main>
  );
}
