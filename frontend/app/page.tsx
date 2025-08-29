import type { Metadata } from "next";
import { LandingHero } from "@/features/landing";

export const metadata: Metadata = {
  title: "TestGo - Your Simple Platform",
  description: "Secure platform for managing your profile and data",
};

export default function Page() {
  return <LandingHero />
}
