#!/bin/sh
SYSTEM_COLOR_FILE=testdata/system-color.csv
COLOR_FILE=testdata/color.csv
OUTPUT_FILE=tables.go
DOC_FILE=colors.md

#############################################################################
# LICENSE
#############################################################################
cat <<EOT > ${OUTPUT_FILE}
// Copyright 2021 Takashi Takizawa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// This file is genarated by the script 'scripts/create-tables.sh'.

package color

EOT

#############################################################################
# const
#############################################################################
cat <<EOT >> ${OUTPUT_FILE}
// Color name
const (
	Default = Color(iota)

	/*
	 * 16 colors for system
	 */

EOT

awk -F, '{
    printf("\t// %s for system color\n", $1);
    printf("\tSystem%s\n", $1);
    printf("\n")
}' < ${SYSTEM_COLOR_FILE} >> ${OUTPUT_FILE}

cat <<EOT >> ${OUTPUT_FILE}
	/*
	 * 256 colors
	 */

EOT

awk -F, '{
    printf("\t// %s: Hex %s, 256-Color Code %s (%s)\n", $1, $2, $3, $4);
    printf("\t%s\n", $1);
    printf("\n")
}' < ${COLOR_FILE} >> ${OUTPUT_FILE}

cat <<EOT >> ${OUTPUT_FILE}
)

EOT

#############################################################################
# var colorNames
#############################################################################

cat <<EOT >> ${OUTPUT_FILE}
var colorNames = []string{
	"Default",

	/*
	 * 16 colors for system
	 */

EOT

awk -F, '{
    printf("\t\"System%s\",\n", $1);
}' < ${SYSTEM_COLOR_FILE} >> ${OUTPUT_FILE}

cat <<EOT >> ${OUTPUT_FILE}

	/*
	 * 256 colors
	 */

EOT

awk -F, '{
    printf("\t\"%s\",\n", $1);
}' < ${COLOR_FILE} >> ${OUTPUT_FILE}

cat <<EOT >> ${OUTPUT_FILE}
}

EOT

#############################################################################
# var colorEscapeSequences
#############################################################################
cat <<EOT >> ${OUTPUT_FILE}
var colorEscapeSequences = [...]string{
	Default: "\x1b[39m",

	/*
	 * 16 colors for system
	 */

EOT

awk -F, '{
    c=sprintf("System%s:", $1)
    printf("\t%-20s \"\\x1b[%sm\",\n", c,$3);
}' < ${SYSTEM_COLOR_FILE} >> ${OUTPUT_FILE}

cat <<EOT >> ${OUTPUT_FILE}

	/*
	 * 256 colors
	 */

EOT

awk -F, '{
    c=sprintf("%s:", $1)
    printf("\t%-12s \"\\x1b[38;5;%sm\",\n", c,$3);
}' < ${COLOR_FILE} >> ${OUTPUT_FILE}

cat <<EOT >> ${OUTPUT_FILE}
}
EOT

#############################################################################
# Document: color tables
#############################################################################

cat <<EOT > ${DOC_FILE}
### System Colors

EOT

awk -F, '{
    printf("- <span style=\"color:%s;\">System%s</span>\n", $2, $1);
}' < ${SYSTEM_COLOR_FILE} >> ${DOC_FILE}

cat <<EOT >> ${DOC_FILE}

### Colors

EOT
awk -F, '{
    printf("- <span style=\"color:%s;\">%s</span>\n", $4, $1);
}' < ${COLOR_FILE} >> ${DOC_FILE}
