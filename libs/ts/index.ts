export enum Level {
    Private = 0, // No Store, No Train
    Personal = 1, // Store Session, Train User
    Global = 2, // Store Perm, Train Base
}

export const HEADER_NAME = "X-PDP-Level";

/**
 * Helper to get just the header object (useful for Axios/Ky clients)
 */
export function getHeader(level: Level): Record<string, string> {
    if (![Level.Private, Level.Personal, Level.Global].includes(level)) {
        level = Level.Private;
    }
    return { [HEADER_NAME]: level.toString() };
}

/**
 * Drop-in wrapper for the native fetch API
 */
export async function fetchWithPDP(
    input: RequestInfo | URL,
    init?: RequestInit,
    level: Level = Level.Private
): Promise<Response> {
    const headers = new Headers(init?.headers);

    // Fail safe: Default to Private if invalid
    if (![Level.Private, Level.Personal, Level.Global].includes(level)) {
        level = Level.Private;
    }

    headers.set(HEADER_NAME, level.toString());

    return fetch(input, {
        ...init,
        headers,
    });
}
