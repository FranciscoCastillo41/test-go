/** @type {import('next').NextConfig} */
const nextConfig = {
    async rewrites() {
      // Use BACKEND_URL (must include /v1) in prod; fall back to local dev
      const target = process.env.BACKEND_URL || "http://localhost:8080/v1";
      return [{ source: "/api/:path*", destination: `${target}/:path*` }];
    },
  };
  
  export default nextConfig;
  