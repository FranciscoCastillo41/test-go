const API_BASE = "/api";

async function handle<T>(res: Response): Promise<T> {
    const txt = await res.text();
    let data: any;
    try {
        data = txt ? JSON.parse(txt) : null;
    } catch {
        data = txt;
    }

    if (!res.ok) {
        const msg =
            (data && (data.error || data.message || data.title)) ||
            `HTTP ${res.status}`;
        throw new Error(msg);
    }
    return data as T;
}

export async function apiGet<T>(path: string, init?: RequestInit): Promise<T> {
    const res = await fetch(`${API_BASE}${path}`, {
        ...init,
        headers: { "Content-Type": "application/json", ...(init?.headers || {}) },
        cache: "no-store",
    });
    return handle<T>(res);
}

export async function apiPost<T>(
    path: string,
    body: unknown,
    init?: RequestInit
): Promise<T> {
    const res = await fetch(`${API_BASE}${path}`, {
        method: "POST",
        body: JSON.stringify(body),
        ...init,
        headers: { "Content-Type": "application/json", ...(init?.headers || {}) },
    });
    return handle<T>(res);
}

export async function apiDelete<T = null>(path: string, init?: RequestInit): Promise<T> {
   const res = await fetch(`${API_BASE}${path}`, {
        method: "DELETE",
        ...init,
        headers: { "Content-Type": "application/json", ...(init?.headers || {}) },
  });
  return handle<T>(res);
}

export async function apiPatch<T>(path: string, body: unknown, init?: RequestInit): Promise<T> {
    const res = await fetch(`${API_BASE}${path}`, {
        method: "PATCH",
        body: JSON.stringify(body),
        ...init,
        headers: { "Content-Type": "application/json", ...(init?.headers || {}) }
    });
    return handle<T>(res);
}