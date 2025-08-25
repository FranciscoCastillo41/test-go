"use client";

import { useState } from "react";
import { useWidget, useDeleteWidget } from "@/features/widgets/hooks";

export default function WidgetFetchById() {
  const [id, setId] = useState("");
  const q = useWidget(id, { enabled: false });
  const del = useDeleteWidget();

  async function onFetch() {
    if (!id.trim()) return;
    await q.refetch();
  }

  async function onDelete() {
    const _id = id.trim();
    if (!_id) return;
    if (!confirm(`Delete widget ${_id}?`)) return;
    await del.mutateAsync(_id);
    // optional: clear view
    setId("");
  }

  return (
    <section style={{ marginTop: 24 }}>
      <h2>Fetch / Delete by ID</h2>
      <input
        placeholder="paste widget id…"
        value={id}
        onChange={(e) => setId(e.target.value)}
        onKeyDown={(e) => e.key === "Enter" && onFetch()}
        style={{ width: "100%" }}
      />
      <div style={{ display: "flex", gap: 8, marginTop: 8 }}>
        <button onClick={onFetch} disabled={q.isFetching}>
          {q.isFetching ? "Fetching..." : "Fetch"}
        </button>
        <button onClick={onDelete} disabled={del.isPending || !id.trim()}>
          {del.isPending ? "Deleting…" : "Delete by ID"}
        </button>
      </div>
      {(q.error as Error)?.message && <p style={{ color: "crimson" }}>Error: {(q.error as Error).message}</p>}
      <pre style={{ marginTop: 16 }}>
        {q.data ? JSON.stringify(q.data, null, 2) : "No widget fetched yet"}
      </pre>
    </section>
  );
}
