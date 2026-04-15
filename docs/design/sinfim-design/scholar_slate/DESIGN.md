# Design System Specification: The Architectural Academic

## 1. Overview & Creative North Star
The "Architectural Academic" is the creative North Star for this design system. In the context of a Learning Management System (LMS), we move away from the cluttered, "utility-first" spreadsheets of the past. Instead, we embrace the quiet authority of a modern research library.

This system rejects the "template" look of standard SaaS dashboards. We achieve a premium, editorial feel through **intentional asymmetry**, high-contrast typographic scales, and **tonal layering**. The layout prioritizes the 1440px desktop experience, treating the screen as a canvas where information breathes. We use wide margins and staggered element placements to guide the learner's eye, moving from "scanning" to "absorbing."

---

## 2. Colors
Our palette is rooted in a deep, authoritative navy and a sophisticated forest green. This establishes a sense of permanence and trust.

### Color Tokens
- **Primary (The Anchor):** `#041632` (Deep Navy). Reserved for high-level navigation and structural identity.
- **Secondary (The Action):** `#2c694e` (Forest Green). Used for primary CTAs and success-pathway progression.
- **Surface (The Canvas):** `#f7f9fb`. A clean, cool-white base that prevents eye strain.

### The "No-Line" Rule
To achieve a high-end feel, **1px solid borders are strictly prohibited for sectioning.** Boundaries must be defined solely through background color shifts.
- **Example:** A sidebar (`primary`) sits flush against the main content area (`surface`), which contains a workspace module (`surface-container-low`). 

### Surface Hierarchy & Nesting
Treat the UI as a series of physical layers. We use the `surface-container` tiers to create "nested" depth:
- **Level 0 (Background):** `surface` (`#f7f9fb`)
- **Level 1 (Sectioning):** `surface-container-low` (`#f2f4f6`)
- **Level 2 (Active Cards):** `surface-container-lowest` (`#ffffff`)

### Signature Textures
To add visual "soul," use subtle gradients on primary actions. A linear gradient from `secondary` (`#2c694e`) to `on-secondary-container` (`#316e52`) provides a tactile, "pressed" quality that flat colors lack.

---

## 3. Typography
We utilize a dual-font strategy to balance editorial sophistication with operational clarity.

- **Display & Headlines (Manrope):** Chosen for its geometric precision. Use `display-lg` to `headline-sm` for page titles and section headers. The wide apertures of Manrope convey a modern, open academic environment.
- **UI & Body (Inter):** The workhorse for the LMS. From `title-md` down to `label-sm`, Inter provides maximum legibility for course content and data-heavy tables.

**Hierarchy Note:** Maintain high contrast between labels and headers. A `label-sm` in uppercase with 5% letter-spacing should be used for secondary metadata to distinguish it from interactive body text.

---

## 4. Elevation & Depth
Depth in this system is achieved through **Tonal Layering** and **Ambient Light**, not heavy drop shadows.

### The Layering Principle
Stack containers to create soft, natural lift. Place a `surface-container-lowest` card on top of a `surface-container-low` section. The subtle shift from `#f2f4f6` to `#ffffff` is enough to signify a new interactive plane.

### Ambient Shadows
When an element must "float" (e.g., a modal or a primary dropdown), use an **extra-diffused shadow**:
- **Y-Offset:** 8px | **Blur:** 24px
- **Color:** `on-surface` (`#191c1e`) at **4% to 8% opacity**. This mimics natural light rather than digital "glow."

### The "Ghost Border" Fallback
If accessibility requires a container boundary, use the **Ghost Border**:
- Token: `outline-variant` (`#c5c6ce`) at **15% opacity**.
- **Rule:** Never use 100% opaque borders; they shatter the editorial flow.

### Glassmorphism
For floating navigation or overlay elements, use a backdrop blur (12px) paired with a semi-transparent `surface` color. This allows the course content to "bleed through," making the layout feel integrated.

---

## 5. Components

### Sidebar (The Navigation Column)
- **Width:** 260px.
- **Style:** Background set to `primary` (`#041632`). 
- **Active State:** Use a glassmorphic treatment—a semi-transparent white overlay with a 4px `lg` radius—to highlight the current page.

### Metric Cards
- **Structure:** No borders. Use `surface-container-lowest` (`#ffffff`) background.
- **Typography:** `headline-md` for the metric, `label-md` for the description.
- **Editorial Touch:** Place a subtle 4px vertical accent bar of `secondary` (`#2c694e`) on the left side to denote importance.

### Structured Data Tables
- **Rule:** **Strictly forbid vertical and horizontal divider lines.**
- **Separation:** Use `surface-container-low` for the header row and alternating `surface` / `surface-container-lowest` for body rows.
- **Hover State:** Shift the row background to `surface-variant` (`#e0e3e5`) for a soft, responsive feel.

### Buttons
- **Primary:** `secondary` base with a subtle gradient. 0.5rem (`lg`) radius.
- **Tertiary (Ghost):** No background or border. Use `on-surface` text with an icon. Upon hover, add a `surface-container-highest` background.

### Status Badges
- Use pill-shaped containers (`full` radius) with high-chroma text on low-chroma backgrounds (e.g., `on-error-container` text on `error-container` background).

---

## 6. Do's and Don'ts

### Do:
- **Embrace White Space:** Use the Spacing Scale to create "breathing room" between modules. 1440px layouts should feel spacious, not packed.
- **Use Tonal Shifting:** Always ask, "Can I define this area with a background color shift instead of a line?"
- **Prioritize the Sidebar:** The 260px navy sidebar is the "anchor" of the user experience; ensure its typography remains high-contrast (`on-primary`).

### Don't:
- **Don't use 100% Black:** Use `on-surface` (`#191c1e`) for text to maintain a sophisticated, soft-ink look.
- **Don't use "Default" Shadows:** Avoid the heavy, blurry shadows common in consumer-grade apps.
- **Don't use Divider Lines:** If content feels cluttered, increase the padding/margin rather than adding a line. Lines are the "clutter of the digital age."