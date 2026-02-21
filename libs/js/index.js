const PDP_HEADER = "X-PDP-Level";

/**
 * fetchWithPDP wraps the native fetch API to automatically inject the PDP header.
 * * @param {string|Request} url - The resource URL.
 * @param {RequestInit} [options={}] - Standard fetch options.
 * @param {number} [pdpLevel=0] - The PDP consent level (0, 1, or 2).
 * @returns {Promise<Response>}
 */
export async function fetchWithPDP(url, options = {}, pdpLevel = 0) {
    if (!Number.isInteger(pdpLevel) || pdpLevel < 0 || pdpLevel > 2) {
        throw new RangeError("PDP level must be an integer between 0 and 2");
    }

    const headers = new Headers(options.headers || {});
    headers.set(PDP_HEADER, pdpLevel.toString());

    return fetch(url, {
        ...options,
        headers,
    });
}
