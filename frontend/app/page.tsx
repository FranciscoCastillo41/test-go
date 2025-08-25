import type { Metadata } from "next";
import WidgetForm from "@/components/widgets/WidgetForm";
import WidgetList from "@/components/widgets/WidgetList";
import WidgetFetchById from "@/components/widgets/WidgetFetchById";

export const metadata: Metadata = {
  title: "Widgets",
  description: "Create, list, and fetch widgets",
};

export default function Page() {
  // Server component (no "use client"): renders client components just fine
  return (
    <main className="stack">
      <h1>Widgets</h1>
      <section>
        <WidgetForm />
      </section>
      <WidgetList />
      <WidgetFetchById />
    </main>
  );
}
