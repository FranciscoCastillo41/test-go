import {
    apiGet,
    apiPost,
    apiDelete,
    apiPatch
} from "../../../lib/api";

import type { Widget } from "../types";

export function fetchWidgets() {
    return apiGet<Widget[]>("/widgets");
}

export function fetchWidget(id: string) {
    return apiGet<Widget>(`/widgets/${id}`);
}

export function createWidget(input: { name: string; price: number }) {
    return apiPost<Widget>("/widgets", input);
}

export function deleteWidget(id: string) {
    return apiDelete<null>(`/widgets/${id}`); // returns null (204)
}

export function deleteAllWidgets() {
    return apiDelete<{ deleted: number }>("/widgets"); // { deleted: n }
}

export function updateWidget(
    id: string,
    patch: Partial<{ name: string; price: number }>
) {
    return apiPatch<Widget>(`/widgets/${id}`, patch);
}