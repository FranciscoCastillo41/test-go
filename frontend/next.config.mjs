/** @type {import('next').NextConfig} */
const nextConfig = {
    // In dev, if NEXT_PUBLIC_API_URL is not set, proxy /api/* → Go backend /v1/*
    async rewrites() {
      if (!process.env.NEXT_PUBLIC_API_URL) {
        return [{ source: "/api/:path*", destination: "http://localhost:8080/v1/:path*" }];
      }
      return [];
    },
  };
  export default nextConfig;
  