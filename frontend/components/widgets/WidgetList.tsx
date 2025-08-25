"use client";

import { useState } from "react";
import { useWidgets, useDeleteWidget, useDeleteAllWidgets, useUpdateWidget } from "@/features/widgets/hooks";
import type { Widget } from "@/features/widgets/types";

function WidgetRow({ w }: { w: Widget }) {
  const delOne = useDeleteWidget();
  const upd = useUpdateWidget();

  const [editing, setEditing] = useState(false);
  const [name, setName] = useState(w.name);
  const [price, setPrice] = useState(w.price.toString()); // string → parse
  const [err, setErr] = useState<string | null>(null);

  async function onSave() {
    setErr(null);
    const patch: Partial<{ name: string; price: number }> = {};
    const trimmed = name.trim();
    const n = parseFloat(price);

    if (trimmed !== w.name) {
      if (!trimmed) return setErr("Name cannot be empty.");
      patch.name = trimmed;
    }
    if (!Number.isNaN(n) && n !== w.price) {
      if (n < 0) return setErr("Price must be ≥ 0.");
      patch.price = n;
    }
    if (Object.keys(patch).length === 0) {
      setEditing(false);
      return;
    }
    try {
      await upd.mutateAsync({ id: w.id, patch });
      setEditing(false);
    } catch (e: any) {
      setErr(e.message ?? "Update failed");
    }
  }

  return (
    <li style={{ display: "flex", alignItems: "center", gap: 8, padding: "8px 0", borderBottom: "1px solid #eee" }}>
      <div style={{ flex: 1 }}>
        {editing ? (
          <>
            <div>
              <label>
                Name{" "}
                <input value={name} onChange={(e) => setName(e.target.value)} />
              </label>
            </div>
            <div>
              <label>
                Price{" "}
                <input
                  type="number"
                  step="0.01"
                  min="0"
                  inputMode="decimal"
                  value={price}
                  onChange={(e) => setPrice(e.target.value)}
                />
              </label>
            </div>
            {err && <p style={{ color: "crimson", marginTop: 4 }}>{err}</p>}
          </>
        ) : (
          <>
            <div><strong>{w.name}</strong> — ${w.price.toFixed(2)}</div>
            <div style={{ fontSize: 12, color: "#666" }}>{w.id}</div>
          </>
        )}
      </div>

      {editing ? (
        <>
          <button onClick={onSave} disabled={upd.isPending}>
            {upd.isPending ? "Saving…" : "Save"}
          </button>
          <button onClick={() => { setEditing(false); setName(w.name); setPrice(w.price.toString()); }} disabled={upd.isPending}>
            Cancel
          </button>
        </>
      ) : (
        <>
          <button onClick={() => setEditing(true)} title="Edit this widget">Edit</button>
          <button onClick={() => delOne.mutate(w.id)} disabled={delOne.isPending} title="Delete this widget">
            {delOne.isPending ? "…" : "Delete"}
          </button>
        </>
      )}
    </li>
  );
}

export default function WidgetList() {
  const q = useWidgets();              // loads list
  const delAll = useDeleteAllWidgets();

  return (
    <section style={{ marginTop: 24 }}>
      <h2>All Widgets</h2>

      <div style={{ display: "flex", gap: 8, marginBottom: 8 }}>
        <button onClick={() => q.refetch()} disabled={q.isFetching}>
          {q.isFetching ? "Loading..." : "Refresh"}
        </button>
        <button
          onClick={async () => {
            if (!q.data?.length) return;
            if (!confirm(`Delete all ${q.data.length} widgets?`)) return;
            await delAll.mutateAsync();
          }}
          disabled={delAll.isPending || !q.data?.length}
        >
          {delAll.isPending ? "Deleting…" : "Delete all"}
        </button>
      </div>

      {(q.error as Error)?.message && <p style={{ color: "crimson" }}>Error: {(q.error as Error).message}</p>}

      {!q.data?.length ? (
        <p>No data yet</p>
      ) : (
        <ul style={{ padding: 0, listStyle: "none" }}>
          {q.data.map((w) => <WidgetRow key={w.id} w={w} />)}
        </ul>
      )}
    </section>
  );
}
