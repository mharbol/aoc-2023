#! /usr/bin/bash

# This is not a very safe script, but it gets the job done for what I need
#
# $1 is the day of the month in XX format

DAY_NUM=$1
DAY_NUM_STRIP=`echo ${DAY_NUM} | sed 's/^0//'`
CURR_DIR=`pwd`
NEW_DAY_DIR=${CURR_DIR}/solution/days/day${DAY_NUM}
SOURCE_FILE=${NEW_DAY_DIR}/day${DAY_NUM}.go
# PART_1="Day${DAY_NUM}Part1"
# PART_2="Day${DAY_NUM}Part2"
SOL_REG_FILE=${CURR_DIR}/solution/solution_registry.go
EXP_REG_FILE=${CURR_DIR}/solution/expected_registry.go

echo "Day number: ${DAY_NUM}"
echo "Trimmed day number: ${DAY_NUM_STRIP}"
echo "Working in ${CURR_DIR}"

# create input txt file
echo "Creating day ${DAY_NUM} input txt file"
touch ${CURR_DIR}/inputs/day_${DAY_NUM}.txt

# create day source file
echo "Creating dir ${NEW_DAY_DIR}/"
mkdir ${NEW_DAY_DIR}

echo "Creating day ${DAY_NUM} solution source file"
cp ${CURR_DIR}/scripts/day_x_template.txt ${SOURCE_FILE}

sed -i "s/DAY/day${DAY_NUM}/g" ${SOURCE_FILE}
# sed -i "s/PART1/${PART_1}/g" ${SOURCE_FILE}
# sed -i "s/PART2/${PART_2}/g" ${SOURCE_FILE}

go fmt ${SOURCE_FILE}

# add day source file to solution registry
echo "Adding Day ${DAY_NUM} to solution registry"
SOL_REG_ENTRY_LINE=`cat ${SOL_REG_FILE} | wc -l`
sed -i "${SOL_REG_ENTRY_LINE}i ${DAY_NUM_STRIP}: {day${DAY_NUM}.Part1, day${DAY_NUM}.Part2}," ${SOL_REG_FILE}

# add import to solution registry
echo "Adding Day ${DAY_NUM} to solution registry import"
IMPORT_LINE_NUM=`sed -n '/^)$/=' ${SOL_REG_FILE}`
sed -i "${IMPORT_LINE_NUM}i \"github.com/mharbol/aoc-2023/solution/days/day${DAY_NUM}\"" ${SOL_REG_FILE}

go fmt ${SOL_REG_FILE}

# add blank solutions to expected registry
echo "Adding Day ${DAY_NUM} to expected registry"
EXP_REG_ENTRY_LINE=`cat ${EXP_REG_FILE} | wc -l`
sed -i "${EXP_REG_ENTRY_LINE}i ${DAY_NUM_STRIP}: {\"\", \"\"}," ${EXP_REG_FILE}
go fmt ${EXP_REG_FILE}

# append README
echo "Writing Day ${DAY_NUM} entry to README.md"
echo "" >> ${CURR_DIR}/README.md
echo "### [Day ${DAY_NUM}](solution/days/day${DAY_NUM}/day${DAY_NUM}.go)" >> ${CURR_DIR}/README.md
echo "TBD" >> ${CURR_DIR}/README.md

exit 0
