"use client";

import { useState } from "react";
import { useCreateWidget } from "@/features/widgets/hooks"

export default function WidgetForm() {
  const [name, setName] = useState("Test Widget");
  const [price, setPrice] = useState("19.99"); // keep string → parse
  const [inputErr, setInputErr] = useState<string | null>(null);
  const create = useCreateWidget();

  async function onSubmit(e: React.FormEvent) {
    e.preventDefault();
    setInputErr(null);
    const n = parseFloat(price);
    if (!name.trim() || Number.isNaN(n)) return setInputErr("Enter a valid name and price.");
    await create.mutateAsync({ name: name.trim(), price: n });
    setName("Test Widget");
    setPrice("19.99");
  }

  return (
    <form onSubmit={onSubmit} className="space-y-2">
      <h2>Create</h2>
      <div>
        <label>Name{" "}
          <input value={name} onChange={(e) => setName(e.target.value)} />
        </label>
      </div>
      <div>
        <label>Price{" "}
          <input
            type="number"
            step="0.01"
            inputMode="decimal"
            min="0"
            value={price}
            onChange={(e) => setPrice(e.target.value)}
          />
        </label>
      </div>
      <button type="submit" disabled={create.isPending}>
        {create.isPending ? "Creating..." : "Create"}
      </button>
      {(inputErr || (create.error as Error)?.message) && (
        <p style={{ color: "crimson" }}>Error: {inputErr ?? (create.error as Error)?.message}</p>
      )}
    </form>
  );
}
