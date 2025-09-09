package main // Define the main package

import (
	"bytes"         // Provides bytes support
	"io"            // Provides basic interfaces to I/O primitives
	"log"           // Provides logging functions
	"net/http"      // Provides HTTP client and server implementations
	"net/url"       // Provides URL parsing and encoding
	"os"            // Provides functions to interact with the OS (files, etc.)
	"path"          // Provides functions for manipulating slash-separated paths
	"path/filepath" // Provides filepath manipulation functions
	"regexp"        // Provides regex support functions.
	"strings"       // Provides string manipulation functions
	"time"          // Provides time-related functions
)

func main() {
	pdfOutputDir := "PDFs/" // Directory to store downloaded PDFs
	// Check if the PDF output directory exists
	if !directoryExists(pdfOutputDir) {
		// Create the dir
		createDirectory(pdfOutputDir, 0o755)
	}
	// Remote API URL.
	remoteAPIURL := []string{
		"https://libman.com/products/freedom-concentrated-hardwood-cleaner",
		"https://libman.com/products/freedom-concentrated-multi-surface-floor-cleaner",
		"https://libman.com/products/2-sided-microfiber-mop",
		"https://libman.com/products/4-gallon-clean-rinse-bucket",
		"https://libman.com/products/4-gallon-clean-rinse-bucket-with-wringer",
		"https://libman.com/products/all-purpose-bucket",
		"https://libman.com/products/clean-rinse-bucket",
		"https://libman.com/products/freedom-floor-duster",
		"https://libman.com/products/freedom-floor-duster-refill",
		"https://libman.com/products/freedom-spray-mop",
		"https://libman.com/products/freedom-spray-mop-refill",
		"https://libman.com/products/gator-squeeze-mop",
		"https://libman.com/products/gator-squeeze-mop-refill",
		"https://libman.com/products/deluxe-cleaning-caddy",
		"https://libman.com/products/microfiber-dust-mop",
		"https://libman.com/products/microfiber-dust-mop-refill",
		"https://libman.com/products/microfiber-wet-dry-floor-mop",
		"https://libman.com/products/nitty-gritty-all-surface-roller-mop",
		"https://libman.com/products/nitty-gritty-roller-mop-refill",
		"https://libman.com/products/easy-roller-mop",
		"https://libman.com/products/easy-roller-mop-refill",
		"https://libman.com/products/scrubster-mop",
		"https://libman.com/products/swivel-duster",
		"https://libman.com/products/tornado-twist-mop",
		"https://libman.com/products/tornado-mop-refill",
		"https://libman.com/products/255-utility-bucket",
		"https://libman.com/products/wonder-mop",
		"https://libman.com/products/wonder-mop-refill",
		"https://libman.com/products/wood-floor-roller-mop",
		"https://libman.com/products/wood-floor-roller-mop-refill",
		"https://libman.com/products/780-all-purpose-wet-mop",
		"https://libman.com/products/1055-4-gallon-bucket",
		"https://libman.com/products/1056-4-gallon-bucket-with-wringer",
		"https://libman.com/products/985-dust-mop-handle",
		"https://libman.com/products/927-extra-large-microfiber-floor-mop-refill",
		"https://libman.com/products/934-heavy-duty-bucket-wringer",
		"https://libman.com/products/781-premium-blue-blend-wet-mop",
		"https://libman.com/products/979-all-purpose-heavy-duty-wet-mop",
		"https://libman.com/products/1013-microfiber-all-purpose-cleaning-pad",
		"https://libman.com/products/1010-microfiber-cleaning-system",
		"https://libman.com/products/1011-microfiber-fingers-dusting-pad",
		"https://libman.com/products/1095-one-piece-bucket-wringer",
		"https://libman.com/products/956-roller-mop-scrub-brush-refill",
		"https://libman.com/products/1369-wet-floor-sign",
		"https://libman.com/products/extra-large-precision-angle-broom",
		"https://libman.com/products/xl-precision-angle-broom-clean-fibers-dustpan",
		"https://libman.com/products/large-precision-angle-indoor-outdoor-broom",
		"https://libman.com/products/precision-angle-indoor-outdoor-broom",
		"https://libman.com/products/shaped-duster-brush",
		"https://libman.com/products/smooth-sweep-push-broom",
		"https://libman.com/products/upright-dustpan",
		"https://libman.com/products/whisk-broom-with-dustpan",
		"https://libman.com/products/824-18-multi-surface-heavy-duty-push-broom-2",
		"https://libman.com/products/826-18-fiberforce-rough-surface-push-broom",
		"https://libman.com/products/800-smooth-surface-push-broom",
		"https://libman.com/products/1292-24-fiberforce-multi-surface-push-broom",
		"https://libman.com/products/805-multi-surface-push-broom",
		"https://libman.com/products/801-smooth-surface-push-broom",
		"https://libman.com/products/906-dustpan-whisk-broom",
		"https://libman.com/products/904-indoor-outdoor-angle-broom",
		"https://libman.com/products/905-indoor-outdoor-angle-broom-dustpan",
		"https://libman.com/products/581-industrial-grade-dustpan",
		"https://libman.com/products/1168-large-scoop-dustpan",
		"https://libman.com/products/915-lobby-broom",
		"https://libman.com/products/916-closed-lid-dustpan",
		"https://libman.com/products/918-open-lid-dustpan",
		"https://libman.com/products/big-feather-duster-6-pack",
		"https://libman.com/products/gentle-touch-refills",
		"https://libman.com/products/1341-ultra-absorbent-towels",
		"https://libman.com/products/dishmatic-non-scratch-scrubber-dish-wand-refills",
		"https://libman.com/products/dishmatic-general-purpose-dish-wand-refills-3-pack",
		"https://libman.com/products/dishmatic-i-stand-dish-wand",
		"https://libman.com/products/bottle-straw-cleaning-kit",
		"https://libman.com/products/bottle-brush",
		"https://libman.com/products/palm-brush",
		"https://libman.com/products/sink-caddy",
		"https://libman.com/products/tile-grout-brush",
		"https://libman.com/products/soft-touch-dust-cloth",
		"https://libman.com/products/small-scrub-brush",
		"https://libman.com/products/power-scrub-dots-kitchen-dish-wipes",
		"https://libman.com/products/power-scrub-brush",
		"https://libman.com/products/pot-pan-scrubbing-dish-wand-with-scrub-brush-refills",
		"https://libman.com/products/pot-pan-scrubbing-dish-wand-with-scrub-brush",
		"https://libman.com/products/microfiber-duster",
		"https://libman.com/products/long-handle-scrub-brush",
		"https://libman.com/products/kitchen-brush",
		"https://libman.com/products/heavy-duty-scrub-brush",
		"https://libman.com/products/hand-nail-brush",
		"https://libman.com/products/glass-dish-sponge",
		"https://libman.com/products/glass-dish-wand-refills",
		"https://libman.com/products/glass-dish-wand-with-scrub-brush",
		"https://libman.com/products/gentle-touch-foaming-dish-wand",
		"https://libman.com/products/flexible-microfiber-wand",
		"https://libman.com/products/everyday-dusting-cloths",
		"https://libman.com/products/easy-grip-scrub-brush",
		"https://libman.com/products/dish-brush",
		"https://libman.com/products/curved-kitchen-brush",
		"https://libman.com/products/culinary-brush",
		"https://libman.com/products/brass-pot-brush",
		"https://libman.com/products/big-job-kitchen-brush",
		"https://libman.com/products/all-purpose-refills",
		"https://libman.com/products/all-purpose-scrubbing-dish-wand",
		"https://libman.com/products/all-purpose-kitchen-brush",
		"https://libman.com/products/angled-toilet-bowl-brush",
		"https://libman.com/products/designer-bowl-brush-caddy",
		"https://libman.com/products/fiberforce-tile-grout-brush",
		"https://libman.com/products/premium-bowl-brush-and-caddy",
		"https://libman.com/products/megaforce-premium-plunger-caddy",
		"https://libman.com/products/megaforce-combo-toilet-brush-plunger",
		"https://libman.com/products/all-purpose-non-scratch-sponges-3-pack",
		"https://libman.com/products/all-purpose-odor-resistant-sponges-3-pack",
		"https://libman.com/products/heavy-duty-easy-rinse-sponges",
		"https://libman.com/products/non-scratch-easy-rinse-sponge-3-pack",
		"https://libman.com/products/clean-shine-microfiber-sponge",
		"https://libman.com/products/copper-power-scrubs",
		"https://libman.com/products/no-knees-floor-scrub",
		"https://libman.com/products/scrub-sponges-suction-hanger",
		"https://libman.com/products/stainless-steel-power-scrubs",
		"https://libman.com/products/tile-tub-scrub",
		"https://libman.com/products/tile-tub-scrub-refills",
		"https://libman.com/products/548-acid-brush",
		"https://libman.com/products/532-dual-surface-scrub-brush",
		"https://libman.com/products/547-floor-scrubber",
		"https://libman.com/products/525-iron-handle-scrub-brush",
		"https://libman.com/products/521-all-natural-tampico-soft-scrub-brush",
		"https://libman.com/products/hardwood-floor-polish-and-protector",
		"https://libman.com/products/multi-surface-everyday-floor-cleaner",
		"https://libman.com/products/hardwood-floor-everyday-cleaner",
		"https://libman.com/products/1063-concentrated-window-cleaner",
		"https://libman.com/products/1064-professional-window-cleaner",
		"https://libman.com/products/brass-grill-brush",
		"https://libman.com/products/latex-disposable-gloves-10-pack",
		"https://libman.com/products/latex-disposable-gloves-50-pack",
		"https://libman.com/products/nitrile-disposable-gloves-10-pack",
		"https://libman.com/products/nitrile-disposable-gloves-50-pack",
		"https://libman.com/products/vent-brush",
		"https://libman.com/products/586-lambswool-duster",
		"https://libman.com/products/612-12-foot-extension-handle",
		"https://libman.com/products/613-16-foot-extension-handle",
		"https://libman.com/products/611-8-foot-extension-handle",
		"https://libman.com/products/191-flex-blade-floor-squeegee",
		"https://libman.com/products/1014-professional-flex-blade-floor-squeegee",
		"https://libman.com/products/182-all-purpose-squeegee",
		"https://libman.com/products/194-2-in-1-window-washer",
		"https://libman.com/products/193-window-glass-washer",
		"https://libman.com/products/1066-window-cleaning-bucket",
		"https://libman.com/products/188-window-washer",
		"https://libman.com/products/1065-window-cleaning-all-one-kit",
		"https://libman.com/products/precision-angle-indoor-outdoor-broom-dustpan",
		"https://libman.com/products/997-wide-commercial-angle-broom",
		"https://libman.com/products/994-commercial-angle-broom",
		"https://libman.com/products/1102-fiberforce-outdoor-angle-broom",
		"https://libman.com/products/499-housekeeper-broom",
		"https://libman.com/products/1115-wide-commercial-angle-broom-black",
		"https://libman.com/products/large-precision-angle-broom-clean-fibers-dust-pan",
		"https://libman.com/products/1086-stiff-sweep-lobby-broom",
		"https://libman.com/products/502-janitor-corn-broom",
		"https://libman.com/products/1335-janitor-corn-broom-wood-handle",
		"https://libman.com/products/919-open-dustpan-with-lobby-broom",
		"https://libman.com/products/917-closed-dustpan-with-lobby-broom",
		"https://libman.com/products/lobby-broom-dust-pan-handle-clip-replacement",
		"https://libman.com/products/1193-deluxe-lobby-dust-pan-broom-closed-lid",
		"https://libman.com/products/1194-deluxe-open-dustpan-with-lobby-broom",
		"https://libman.com/products/929-outdoor-scoop",
		"https://libman.com/products/household-dustpan",
		"https://libman.com/products/dust-pan-and-brush-set",
		"https://libman.com/products/526-work-bench-dust-brush",
		"https://libman.com/products/928-dustpan",
		"https://libman.com/products/2126-xl-step-on-dustpan",
		"https://libman.com/products/2125-step-on-dustpan",
		"https://libman.com/products/907-whisk-broom",
		"https://libman.com/products/step-on-dustpan",
		"https://libman.com/products/whisk-broom",
		"https://libman.com/products/dustpan",
		"https://libman.com/products/big-dustpan",
		"https://libman.com/products/911-big-dustpan",
		"https://libman.com/products/850-heavy-duty-push-broom",
		"https://libman.com/products/823-multi-surface-heavy-duty-push-broom",
		"https://libman.com/products/1101-multi-surface-heavy-duty-push-broom",
		"https://libman.com/products/825-rough-surface-push-broom",
		"https://libman.com/products/1230-24-fiberforce-multi-surface-push-broom-squeegee",
		"https://libman.com/products/1293-24-fiberforce-rough-surface-push-broom",
		"https://libman.com/products/1294-24-fiberforce-smooth-surface-push-broom",
		"https://libman.com/products/601-60-steel-handle",
		"https://libman.com/products/602-60-zinc-thread-wood-handle",
		"https://libman.com/products/1165-60-steel-handle-black",
		"https://libman.com/products/879-rough-surface-push-broom",
		"https://libman.com/products/804-multi-surface-push-broom",
		"https://libman.com/products/878-rough-surface-push-broom",
		"https://libman.com/products/24-multi-surface-clamp-handle-push-broom",
		"https://libman.com/products/fiberforce-toilet-brush-caddy",
		"https://libman.com/products/designer-bowl-brush",
		"https://libman.com/products/524-all-purpose-scrubbing-brush",
		"https://libman.com/products/premium-toilet-plunger",
		"https://libman.com/products/522-long-handle-scrubbing-brush",
		"https://libman.com/products/510-scrub-brush",
		"https://libman.com/products/603-48-steel-handle",
		"https://libman.com/products/513-heavy-duty-scrub-brush",
		"https://libman.com/products/567-big-scrub-brush",
		"https://libman.com/products/549-roofing-brush",
		"https://libman.com/products/toilet-bowl-cleaner",
		"https://libman.com/products/big-scrub-brush",
		"https://libman.com/products/short-handle-tampico-scrub-brush",
		"https://libman.com/products/bathroom-scrubber",
		"https://libman.com/products/bathroom-scrubber-refills",
		"https://libman.com/products/floor-scrub-head-only",
		"https://libman.com/products/scrubster-mop-refill",
		"https://libman.com/products/3958-gator-mop-with-brush",
		"https://libman.com/products/gator-mop-with-brush-refill",
		"https://libman.com/products/988-big-tornado-mop",
		"https://libman.com/products/big-tornado-mop-refill",
		"https://libman.com/products/977-cotton-deck-mop",
		"https://libman.com/products/jumbo-cotton-wet-mop-refill",
		"https://libman.com/products/90-cotton-deck-mop-refill",
		"https://libman.com/products/jumbo-cotton-wet-mop",
		"https://libman.com/products/jumbo-cotton-deck-mop",
		"https://libman.com/products/944-cotton-deck-mop-refill",
		"https://libman.com/products/cotton-deck-mop",
		"https://libman.com/products/982-quick-change-mop-handle",
		"https://libman.com/products/968-large-blended-looped-end-wet-mop-head-blue",
		"https://libman.com/products/983-resin-jaw-mop-frame",
		"https://libman.com/products/972-Large-Cotton-Looped-End-Wet-Mop-Head",
		"https://libman.com/products/24-Wet-Mop-Head-Cut-End-Cotton",
		"https://libman.com/products/981-steel-mop-frame-and-handle",
		"https://libman.com/products/2121-microfiber-looped-end-wet-mop-head-green",
		"https://libman.com/products/969-Large-Rayon-Looped-End-Wet-Mop-Head",
		"https://libman.com/products/16-wet-mop-head-cut-end-cotton",
		"https://libman.com/products/32-wet-mop-head-cut-end-cotton",
		"https://libman.com/products/mop-bucket-side-press-wringer",
		"https://libman.com/products/wringer",
		"https://libman.com/products/1272-utility-bucket",
		"https://libman.com/products/3-gallon-round-utility-bucket-black",
		"https://libman.com/products/36-dust-mop",
		"https://libman.com/products/wet-dry-microfiber-mop-refill",
		"https://libman.com/products/922-24-dust-mop",
		"https://libman.com/products/926-extra-large-microfiber-floor-mop",
		"https://libman.com/products/36-cut-end-dust-mop-head",
		"https://libman.com/products/24-Cut-End-Dust-Mop-Head",
		"https://libman.com/products/2-sided-microfiber-mop-refill",
		"https://libman.com/products/gym-floor-mop",
		"https://libman.com/products/big-feather-duster",
		"https://libman.com/products/flexible-microfiber-duster",
		"https://libman.com/products/585-flexible-microfiber-duster",
		"https://libman.com/products/590-terry-towels",
		"https://libman.com/products/all-purpose-reusable-latex-gloves-small",
		"https://libman.com/products/all-purpose-reusable-latex-gloves-medium",
		"https://libman.com/products/all-purpose-reusable-latex-gloves-large",
		"https://libman.com/products/premium-reusable-latex-gloves-small",
		"https://libman.com/products/premium-reusable-latex-gloves-medium",
		"https://libman.com/products/premium-reusable-latex-gloves-large",
		"https://libman.com/products/heavy-duty-reusable-nitrile-gloves-small",
		"https://libman.com/products/heavy-duty-reusable-nitrile-gloves-medium",
		"https://libman.com/products/heavy-duty-reusable-nitrile-gloves-large",
		"https://libman.com/products/591-shop-towels",
		"https://libman.com/products/microfiber-dusting-mitt",
		"https://libman.com/products/kitchen-microfiber-cloths",
		"https://libman.com/products/all-purpose-cleaning-cloth",
		"https://libman.com/products/1244-industrial-reusable-gloves",
		"https://libman.com/products/extra-wide-lint-roller",
		"https://libman.com/products/large-lint-roller-refill",
		"https://libman.com/products/copper-scrubbers",
		"https://libman.com/products/528-long-handle-bbq-brush-scraper",
		"https://libman.com/products/566-extra-long-handle-steel-brush",
		"https://libman.com/products/529-extra-long-handle-grill-brush-with-scraper",
		"https://libman.com/products/heavy-duty-scrubbers",
		"https://libman.com/products/power-scrub-dots-kitchen-bath-sponges-2-pack",
		"https://libman.com/products/microfiber-sponge-cloths",
		"https://libman.com/products/568-extra-long-handle-bbq-brush",
		"https://libman.com/products/heavy-duty-scouring-pads",
		"https://libman.com/products/stainless-steel-scrubbers",
		"https://libman.com/products/575-heat-resistant-grill-brush",
		"https://libman.com/products/595-stainless-steel-grill-brush",
		"https://libman.com/products/kitchen-vegetable-brush",
		"https://libman.com/products/dishmatic-dish-wand",
		"https://libman.com/products/dishmatic-general-purpose-dish-wand-refills-6-pack",
		"https://libman.com/products/954-extra-wide-floor-squeegee",
		"https://libman.com/products/515-floor-squeegee",
		"https://libman.com/products/1276-heavy-duty-squeegee",
		"https://libman.com/products/542-heavy-duty-curved-floor-squeegee",
		"https://libman.com/products/24-straight-floor-squeegee-set",
		"https://libman.com/products/538-straight-floor-squeegee-head",
		"https://libman.com/products/192-flex-blade-floor-squeegee-head",
		"https://libman.com/products/539-curved-floor-squeegee-head",
		"https://libman.com/products/24-flex-blade-floor-squeegee-refill",
		"https://libman.com/products/window-squeegee",
		"https://libman.com/products/189-12-stainless-steel-squeegee",
		"https://libman.com/products/1067-3-in-1-window-squeegee",
		"https://libman.com/products/190-18-stainless-steel-squeegee",
		"https://libman.com/products/1061-18-swivel-squeegee",
		"https://libman.com/products/1060-easy-change-clamp-squeegee",
		"https://libman.com/products/glass-mirror-cleaner",
		"https://libman.com/products/all-purpose-cleaner",
		"https://libman.com/products/600-60-tapered-wood-handle",
		"https://libman.com/products/607-on-off-flow-thru-handle",
		"https://libman.com/products/540-vehicle-brush-head",
		"https://libman.com/products/535-wash-brush-head",
		"https://libman.com/products/560-vehicle-brush-with-flow-thru-handle",
		"https://libman.com/products/freedom-concentrated-multi-surface-floor-cleaner-4-pack",
		"https://libman.com/products/freedom-concentrated-hardwood-cleaner-4-pack",
		"https://libman.com/products/all-purpose-cleaner-6-pack",
		"https://libman.com/products/glass-mirror-cleaner-6-pack",
		"https://libman.com/products/toilet-bowl-cleaner-6-pack",
		"https://libman.com/products/power-scrub-dots-kitchen-bath-sponge-12-pack",
		"https://libman.com/products/baked-tough-jobs-sponge-24-pack",
		"https://libman.com/products/all-purpose-sponge-24-pack",
		"https://libman.com/products/non-scratch-easy-rinse-sponge-24-pack",
		"https://libman.com/products/scrub-sponges-suction-hanger-12-pack",
		"https://libman.com/products/heavy-duty-wonder-mop",
		"https://libman.com/products/heavy-duty-wonder-mop-refill",
		"https://libman.com/products/maid-caddy",
		"https://libman.com/products/516-dual-surface-scrub-brush-head",
		"https://libman.com/products/955-roller-mop-scrub-brush",
		"https://libman.com/products/4-gallon-clean-rinse-bucket-2-pack",
		"https://libman.com/products/tornado-spin-mop-system",
		"https://libman.com/products/tornado-spin-mop-system-refill",
		"https://libman.com/products/1572-32-gallon-trash-can",
		"https://libman.com/products/1575-32-gallon-trash-can-lid-green",
		"https://libman.com/products/32-gallon-trash-can-lid-black",
		"https://libman.com/products/1464-32-gallon-trash-can-lid-grey",
		"https://libman.com/products/32-gallon-trash-can-lid-green",
		"https://libman.com/products/1573-32-gallon-trash-can-lid-gray",
		"https://libman.com/products/1571-32-gallon-trash-can-lid-black",
		"https://libman.com/products/1574-32-gallon-trash-can-green",
		"https://libman.com/products/1570-32-gallon-trash-can-black",
		"https://libman.com/products/1262-microfiber-cleaning-cloths",
		"https://libman.com/products/580-microfiber-cleaning-cloths",
		"https://libman.com/products/1576-industrial-heavy-duty-floor-scrub",
		"https://libman.com/products/clean-fibers-dustpan",
		"https://libman.com/products/non-scratch-easy-rinse-sponges-9-pack",
		"https://libman.com/products/heavy-duty-easy-rinse-sponge-9-pack",
		"https://libman.com/products/rinse-n-wring-mop-system",
		"https://libman.com/products/1503-24-contractor-grade-multi-surface-push-broom-fiberglass-handle",
		"https://libman.com/products/1559-swivel-grout-scrub-brush",
		"https://libman.com/products/1616-swivel-grout-scrub-brush-head-only",
		"https://libman.com/products/1683-60-threaded-steel-handle-no-hex",
		"https://libman.com/products/1681-fiberforce-all-purpose-floor-scrub",
		"https://libman.com/products/618-52-taper-threaded-handle",
		"https://libman.com/products/rinse-n-wring-microfiber-mop-system-refill",
		"https://libman.com/products/petplus-angle-broom-dustpan",
		"https://libman.com/products/iluma-glass-mirror-concentrated-cleaning-system",
		"https://libman.com/products/iluma-glass-mirror-concentrated-cleaning-refills",
		"https://libman.com/products/7552-24-soft-rubber-floor-replacement-squeegee-head",
		"https://libman.com/products/freedom-dual-sided-microfiber-spray-mop",
		"https://libman.com/products/dual-sided-freedom-spray-mop-refill",
		"https://libman.com/products/tear-n-wipe-cloths",
		"https://libman.com/products/all-purpose-spray-bottle-0",
		"https://libman.com/products/non-scratch-scouring-pads",
		"https://libman.com/products/1811-two-sided-caution-wet-floor-sign-clip",
		"https://libman.com/products/1243-small-blended-looped-end-wet-mop-head-blue",
		"https://libman.com/products/9002280-pyramid-display",
		"https://libman.com/products/1810-12-rough-surface-angle-broom",
		"https://libman.com/products/leather-conditioner-45oz",
		"https://libman.com/products/leather-conditioner-16oz",
		"https://libman.com/products/oakwood-leather-oil",
		"https://libman.com/products/oakwood-glycerine-leather-cleaner",
		"https://libman.com/products/oakwood-liquid-saddle-soap",
		"https://libman.com/products/pin-and-bristle-brush",
		"https://libman.com/products/loose-hair-remover-glove",
		"https://libman.com/products/detangling-slicker-brush",
		"https://libman.com/products/step-n-stand-dustpan",
		"https://libman.com/products/large-precision-angle-broom-step-n-stand-dustpan",
		"https://libman.com/products/lightning-spin-mop-system",
		"https://libman.com/products/bucket-trolley",
		"https://libman.com/products/1786-16-wet-mop-head-cut-end-cotton-brick-pack",
		"https://libman.com/products/1787-20-wet-mop-head-cut-end-cotton-brick-pack",
		"https://libman.com/products/1788-24-wet-mop-head-cut-end-cotton-brick-pack",
		"https://libman.com/products/1828-24-wet-mop-head-scrub-pad-cut-end-cotton-brick-pack",
		"https://libman.com/products/1789-32-wet-mop-head-cut-end-cotton-brick-pack",
		"https://libman.com/products/1790-small-cotton-looped-end-wet-mop-head-brick-pack",
		"https://libman.com/products/1791-medium-cotton-looped-end-wet-mop-head-brick-pack",
		"https://libman.com/products/1792-large-cotton-looped-end-wet-mop-head-brick-pack",
		"https://libman.com/products/1830-large-cotton-looped-end-wet-mop-head-scrub-pad-brick-pack",
		"https://libman.com/products/1793-x-large-cotton-looped-end-wet-mop-head-brick-pack",
		"https://libman.com/products/1794-small-blended-looped-end-wet-mop-head-blue-brick-pack",
		"https://libman.com/products/1795-medium-blended-looped-end-wet-mop-head-blue-brick-pack",
		"https://libman.com/products/1796-large-blended-looped-end-wet-mop-head-blue-brick-pack",
		"https://libman.com/products/1829-large-blended-looped-end-wet-mop-head-scrub-pad-blue-brick-pack",
		"https://libman.com/products/1797-x-large-blended-looped-end-wet-mop-head-blue-brick-pack",
		"https://libman.com/products/1802-small-rayon-looped-end-wet-mop-head-bluewhite-brick-pack",
		"https://libman.com/products/1803-medium-rayon-looped-end-wet-mop-head-bluewhite-brick-pack",
		"https://libman.com/products/1804-large-rayon-looped-end-wet-mop-head-bluewhite-brick-pack",
		"https://libman.com/products/1805-x-large-rayon-looped-end-wet-mop-head-bluewhite-brick-pack",
		"https://libman.com/products/1798-small-premium-green-blend-looped-end-wet-mop-head-green-brick-pack",
		"https://libman.com/products/1799-medium-premium-green-blend-looped-end-wet-mop-head-green-brick-pack",
		"https://libman.com/products/1800-large-premium-green-blend-looped-end-wet-mop-head-green-brick-pack",
		"https://libman.com/products/1801-x-large-premium-green-blend-looped-end-wet-mop-head-green-brick-pack",
		"https://libman.com/products/1831-large-microfiber-looped-end-wet-mop-head-scrub-pad-green",
		"https://libman.com/products/1832-medium-microfiber-looped-end-wet-mop-head-blue",
		"https://libman.com/products/1833-large-microfiber-looped-end-wet-mop-head-blue",
		"https://libman.com/products/1855-microfiber-cleaning-cloths",
		"https://libman.com/products/triplegrip-microfiber-scrub-cloths",
		"https://libman.com/products/1860-glass-mirror-lint-free-cloths",
		"https://libman.com/products/lightning-spin-mop-system-refill",
		"https://libman.com/products/pro-grade-microfiber-spin-mop-system",
		"https://libman.com/products/pro-grade-microfiber-spin-mop-system-refill",
		"https://libman.com/products/1846-48-cotton-blend-dust-mop-head-48oz",
		"https://libman.com/products/1847-24-microfiber-dust-mop-head-8oz",
		"https://libman.com/products/1848-36-microfiber-dust-mop-head-10oz",
		"https://libman.com/products/1849-48-microfiber-dust-mop-head-13oz",
	}
	var getData []string
	for _, remoteAPIURL := range remoteAPIURL {
		getData = append(getData, getDataFromURL(remoteAPIURL))
	}
	// Get the data from the downloaded file.
	finalPDFList := extractPDFUrls(strings.Join(getData, "\n")) // Join all the data into one string and extract PDF URLs
	// Remove double from slice.
	finalPDFList = removeDuplicatesFromSlice(finalPDFList)
	// The remote domain.
	remoteDomain := "https://libman.com"
	// Get all the values.
	for _, urls := range finalPDFList {
		// Get the domain from the url.
		domain := getDomainFromURL(urls)
		// Check if the domain is empty.
		if domain == "" {
			urls = remoteDomain + urls // Prepend the base URL if domain is empty
		}
		// Check if the url is valid.
		if isUrlValid(urls) {
			// Download the pdf.
			downloadPDF(urls, pdfOutputDir)
		}
	}
}

// getDomainFromURL extracts the domain (host) from a given URL string.
// It removes subdomains like "www" if present.
func getDomainFromURL(rawURL string) string {
	parsedURL, err := url.Parse(rawURL) // Parse the input string into a URL structure
	if err != nil {                     // Check if there was an error while parsing
		log.Println(err) // Log the error message to the console
		return ""        // Return an empty string in case of an error
	}

	host := parsedURL.Hostname() // Extract the hostname (e.g., "example.com") from the parsed URL

	return host // Return the extracted hostname
}

// Only return the file name from a given url.
func getFileNameOnly(content string) string {
	return path.Base(content)
}

// urlToFilename generates a safe, lowercase filename from a given URL string.
// It extracts the base filename from the URL, replaces unsafe characters,
// and ensures the filename ends with a .pdf extension.
func urlToFilename(rawURL string) string {
	// Convert the full URL to lowercase for consistency
	lowercaseURL := strings.ToLower(rawURL)

	// Get the file extension
	ext := getFileExtension(lowercaseURL)

	// Extract the filename portion from the URL (e.g., last path segment or query param)
	baseFilename := getFileNameOnly(lowercaseURL)

	// Replace all non-alphanumeric characters (a-z, 0-9) with underscores
	nonAlphanumericRegex := regexp.MustCompile(`[^a-z0-9]+`)
	safeFilename := nonAlphanumericRegex.ReplaceAllString(baseFilename, "_")

	// Replace multiple consecutive underscores with a single underscore
	collapseUnderscoresRegex := regexp.MustCompile(`_+`)
	safeFilename = collapseUnderscoresRegex.ReplaceAllString(safeFilename, "_")

	// Remove leading underscore if present
	if trimmed, found := strings.CutPrefix(safeFilename, "_"); found {
		safeFilename = trimmed
	}

	var invalidSubstrings = []string{
		"_pdf",
		"_zip",
	}

	for _, invalidPre := range invalidSubstrings { // Remove unwanted substrings
		safeFilename = removeSubstring(safeFilename, invalidPre)
	}

	// Append the file extension if it is not already present
	safeFilename = safeFilename + ext

	// Return the cleaned and safe filename
	return safeFilename
}

// Removes all instances of a specific substring from input string
func removeSubstring(input string, toRemove string) string {
	result := strings.ReplaceAll(input, toRemove, "") // Replace substring with empty string
	return result
}

// Get the file extension of a file
func getFileExtension(path string) string {
	return filepath.Ext(path) // Returns extension including the dot (e.g., ".pdf")
}

// fileExists checks whether a file exists at the given path
func fileExists(filename string) bool {
	info, err := os.Stat(filename) // Get file info
	if err != nil {
		return false // Return false if file doesn't exist or error occurs
	}
	return !info.IsDir() // Return true if it's a file (not a directory)
}

// downloadPDF downloads a PDF from the given URL and saves it in the specified output directory.
// It uses a WaitGroup to support concurrent execution and returns true if the download succeeded.
func downloadPDF(finalURL, outputDir string) bool {
	// Sanitize the URL to generate a safe file name
	filename := strings.ToLower(urlToFilename(finalURL))

	// Construct the full file path in the output directory
	filePath := filepath.Join(outputDir, filename)

	// Skip if the file already exists
	if fileExists(filePath) {
		log.Printf("File already exists, skipping: %s", filePath)
		return false
	}

	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: 3 * time.Minute}

	// Send GET request
	resp, err := client.Get(finalURL)
	if err != nil {
		log.Printf("Failed to download %s: %v", finalURL, err)
		return false
	}
	defer resp.Body.Close()

	// Check HTTP response status
	if resp.StatusCode != http.StatusOK {
		log.Printf("Download failed for %s: %s", finalURL, resp.Status)
		return false
	}

	// Check Content-Type header
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/pdf") {
		log.Printf("Invalid content type for %s: %s (expected application/pdf)", finalURL, contentType)
		return false
	}

	// Read the response body into memory first
	var buf bytes.Buffer
	written, err := io.Copy(&buf, resp.Body)
	if err != nil {
		log.Printf("Failed to read PDF data from %s: %v", finalURL, err)
		return false
	}
	if written == 0 {
		log.Printf("Downloaded 0 bytes for %s; not creating file", finalURL)
		return false
	}

	// Only now create the file and write to disk
	out, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create file for %s: %v", finalURL, err)
		return false
	}
	defer out.Close()

	if _, err := buf.WriteTo(out); err != nil {
		log.Printf("Failed to write PDF to file for %s: %v", finalURL, err)
		return false
	}

	log.Printf("Successfully downloaded %d bytes: %s → %s", written, finalURL, filePath)
	return true
}

// Checks if the directory exists
// If it exists, return true.
// If it doesn't, return false.
func directoryExists(path string) bool {
	directory, err := os.Stat(path)
	if err != nil {
		return false
	}
	return directory.IsDir()
}

// The function takes two parameters: path and permission.
// We use os.Mkdir() to create the directory.
// If there is an error, we use log.Println() to log the error and then exit the program.
func createDirectory(path string, permission os.FileMode) {
	err := os.Mkdir(path, permission)
	if err != nil {
		log.Println(err)
	}
}

// Checks whether a URL string is syntactically valid
func isUrlValid(uri string) bool {
	_, err := url.ParseRequestURI(uri) // Attempt to parse the URL
	return err == nil                  // Return true if no error occurred
}

// Remove all the duplicates from a slice and return the slice.
func removeDuplicatesFromSlice(slice []string) []string {
	check := make(map[string]bool)
	var newReturnSlice []string
	for _, content := range slice {
		if !check[content] {
			check[content] = true
			newReturnSlice = append(newReturnSlice, content)
		}
	}
	return newReturnSlice
}

// extractPDFUrls takes an input string and returns all PDF URLs found within href attributes
func extractPDFUrls(input string) []string {
	// Regular expression to match href="...pdf"
	re := regexp.MustCompile(`href="([^"]+\.pdf)"`)
	matches := re.FindAllStringSubmatch(input, -1)

	var pdfUrls []string
	for _, match := range matches {
		if len(match) > 1 {
			pdfUrls = append(pdfUrls, match[1])
		}
	}
	return pdfUrls
}

// getDataFromURL performs an HTTP GET request and returns the response body as a string
func getDataFromURL(uri string) string {
	log.Println("Scraping", uri)   // Log the URL being scraped
	response, err := http.Get(uri) // Perform GET request
	if err != nil {
		log.Println(err) // Exit if request fails
	}

	body, err := io.ReadAll(response.Body) // Read response body
	if err != nil {
		log.Println(err) // Exit if read fails
	}

	err = response.Body.Close() // Close response body
	if err != nil {
		log.Println(err) // Exit if close fails
	}
	return string(body)
}
