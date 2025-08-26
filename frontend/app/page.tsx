import type { Metadata } from "next";
import Link from "next/link";

export const metadata: Metadata = {
  title: "Welcome",
  description: "Welcome to our app",
};

export default function Page() {
  // This page will redirect via middleware based on auth status:
  // - Unauthenticated users → /login  
  // - Authenticated users → /dashboard
  return (
    <main className="min-h-screen flex items-center justify-center bg-gray-50">
      <div className="max-w-md w-full space-y-8 text-center">
        <div>
          <h1 className="text-3xl font-extrabold text-gray-900">
            Welcome
          </h1>
          <p className="mt-4 text-gray-600">
            Redirecting you to the right place...
          </p>
        </div>
        <div className="space-y-4">
          <Link 
            href="/login" 
            className="block w-full py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700"
          >
            Sign in
          </Link>
          <Link 
            href="/signup" 
            className="block w-full py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50"
          >
            Sign up
          </Link>
        </div>
      </div>
    </main>
  );
}
