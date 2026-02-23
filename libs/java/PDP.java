package pdp;

import java.util.Collections;
import java.util.Map;

public class PDP {
    public static final int LEVEL_PRIVATE = 0; // No Store, No Train
    public static final int LEVEL_PERSONAL = 1; // Store Session, Train User
    public static final int LEVEL_GLOBAL = 2;   // Store Perm, Train Base

    public static final String HEADER_NAME = "X-PDP-Level";

    /**
     * Returns an immutable map containing the PDP header.
     * Defaults to LEVEL_PRIVATE (0) if the input is invalid.
     */
    public static Map<String, String> getHeader(int level) {
        if (level < LEVEL_PRIVATE || level > LEVEL_GLOBAL) {
            level = LEVEL_PRIVATE;
        }
        return Collections.singletonMap(HEADER_NAME, String.valueOf(level));
    }
}
