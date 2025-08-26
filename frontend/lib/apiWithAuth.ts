import { createClient } from '@/lib/supabase/client'

const API_BASE = "/api"

const supabase = createClient()

async function getAuthHeaders(): Promise<Record<string, string>> {
  const { data: { session } } = await supabase.auth.getSession()
  
  if (session?.access_token) {
    return {
      'Authorization': `Bearer ${session.access_token}`,
      'Content-Type': 'application/json'
    }
  }
  
  return { 'Content-Type': 'application/json' }
}

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
    const authHeaders = await getAuthHeaders()
    
    const res = await fetch(`${API_BASE}${path}`, {
        ...init,
        headers: { ...authHeaders, ...(init?.headers || {}) },
        cache: "no-store",
    });
    return handle<T>(res);
}

export async function apiPost<T>(
    path: string,
    body: unknown,
    init?: RequestInit
): Promise<T> {
    const authHeaders = await getAuthHeaders()
    
    const res = await fetch(`${API_BASE}${path}`, {
        method: "POST",
        body: JSON.stringify(body),
        ...init,
        headers: { ...authHeaders, ...(init?.headers || {}) },
    });
    return handle<T>(res);
}

export async function apiDelete<T = null>(path: string, init?: RequestInit): Promise<T> {
    const authHeaders = await getAuthHeaders()
    
   const res = await fetch(`${API_BASE}${path}`, {
        method: "DELETE",
        ...init,
        headers: { ...authHeaders, ...(init?.headers || {}) },
  });
  return handle<T>(res);
}

export async function apiPatch<T>(path: string, body: unknown, init?: RequestInit): Promise<T> {
    const authHeaders = await getAuthHeaders()
    
    const res = await fetch(`${API_BASE}${path}`, {
        method: "PATCH",
        body: JSON.stringify(body),
        ...init,
        headers: { ...authHeaders, ...(init?.headers || {}) }
    });
    return handle<T>(res);
}

export async function apiPut<T>(path: string, body: unknown, init?: RequestInit): Promise<T> {
    const authHeaders = await getAuthHeaders()
    
    const res = await fetch(`${API_BASE}${path}`, {
        method: "PUT",
        body: JSON.stringify(body),
        ...init,
        headers: { ...authHeaders, ...(init?.headers || {}) }
    });
    return handle<T>(res);
}