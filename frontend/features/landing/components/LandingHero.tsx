import Link from 'next/link'

export default function LandingHero() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-indigo-50 via-white to-cyan-50 flex items-center justify-center px-4">
      <div className="max-w-4xl mx-auto text-center">
        {/* Hero Content */}
        <div className="space-y-8">
          {/* Main Heading */}
          <div className="space-y-4">
            <h1 className="text-4xl md:text-6xl font-bold text-gray-900 tracking-tight">
              Welcome to{' '}
              <span className="bg-gradient-to-r from-indigo-600 to-cyan-600 bg-clip-text text-transparent">
                TestGo
              </span>
            </h1>
            <p className="text-xl md:text-2xl text-gray-600 max-w-2xl mx-auto">
              Your simple, secure platform for managing your profile and data.
            </p>
          </div>

          {/* CTA Buttons */}
          <div className="flex flex-col sm:flex-row gap-4 justify-center items-center">
            <Link
              href="/signup"
              className="w-full sm:w-auto bg-indigo-600 hover:bg-indigo-700 text-white font-semibold py-3 px-8 rounded-lg transition-colors duration-200 shadow-lg hover:shadow-xl"
            >
              Get Started
            </Link>
            <Link
              href="/login"
              className="w-full sm:w-auto bg-white hover:bg-gray-50 text-gray-900 font-semibold py-3 px-8 rounded-lg border border-gray-300 transition-colors duration-200 shadow-sm hover:shadow-md"
            >
              Sign In
            </Link>
          </div>

          {/* Simple Feature Highlight */}
          <div className="pt-8">
            <p className="text-sm text-gray-500">
              ✨ Secure authentication • 🔒 Profile management • ⚡ Fast & simple
            </p>
          </div>
        </div>
      </div>
    </div>
  )
}