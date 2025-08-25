"use client";

import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import {
    fetchWidgets,
    fetchWidget,
    createWidget,
    deleteWidget,
    deleteAllWidgets,
    updateWidget
} from "./api";
import type { Widget } from "./types";

export function useWidgets(opts?: { enabled?: boolean }) {
    return useQuery<Widget[]>({
        queryKey: ["widgets"],
        queryFn: fetchWidgets,
        enabled: opts?.enabled ?? true,
        staleTime: 10_000,
    });
}

export function useWidget(id: string, opts?: { enabled?: boolean }) {
    return useQuery<Widget>({
        queryKey: ["widget", id],
        queryFn: () => fetchWidget(id),
        enabled: (opts?.enabled ?? true) && !!id,
        staleTime: 10_000,
    });
}

export function useCreateWidget() {
    const qc = useQueryClient();
    return useMutation({
        mutationFn: (input: { name: string; price: number }) => createWidget(input),
        onSuccess: () => {
            qc.invalidateQueries({ queryKey: ["widgets"] });
        },
    });
}

export function useDeleteWidget() {
    const qc = useQueryClient();
    return useMutation({
        mutationFn: (id: string) => deleteWidget(id),
        onSuccess: () => qc.invalidateQueries({ queryKey: ["widgets"] }),
    })
}

export function useDeleteAllWidgets() {
    const qc = useQueryClient();
    return useMutation({
        mutationFn: () => deleteAllWidgets(),
        onSuccess: () => qc.invalidateQueries({ queryKey: ["widgets"] }),
    });
}

export function useUpdateWidget() {
    const qc = useQueryClient();
    return useMutation({
        mutationFn: ({ id, patch }: { id: string; patch: Partial<Pick<Widget, "name" | "price">> }) =>
            updateWidget(id, patch),
        onSuccess: (data) => {
            qc.invalidateQueries({ queryKey: ["widgets"] });
            qc.invalidateQueries({ queryKey: ["widget", data.id] });
        },
    });
}