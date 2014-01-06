#!/bin/bash

command=$1

help () {
  title "Usage:"
  row "    mao [module] [command]"
  title "Commands:"
  row  "    help         Show this help and exit."
  echo ""
}

run () {
  if [ ! -f "$1" ];
  then
      echo "
    Oops, \"$1\" doesn't exist.
"
      exit 1
  fi

  local modulePath=$(echo "$1" | sed 's/\([^\/]*\)\.go$/mao_\1.go/')
  local errorOutputPath="/tmp/mao-errors-"$[ ( $RANDOM % 99999 )  + 1 ]
  rm -f "$modulePath"
  local source=$(cat $1)
  local module="package main
import . \"github.com/azer/mao\""

  local bddStartsAt=$(echo "$source" | grep -n "Desc(" -m 1 |cut -f1 -d:)

  if [ -z "$bddStartsAt" ]; then
    bddStartsAt=$(echo "$source" | grep -n "Describe(" -m 1 |cut -f1 -d:)
  fi

  if [ -z "$bddStartsAt" ]; then
    printError "Unable to find any 'Desc' or 'Describe' call in '$1'"
  fi

  local wrapped=$(echo "$source" | sed "${bddStartsAt}i func main() {
  ")

  module="$module
  $wrapped
  PrintTestSummary()
  }"

  echo "$module" >> "$modulePath"

  result=$(MAO_LINENO_START=3 go run "$modulePath" | tee /dev/tty)
  rm -f "$modulePath"

  printf "\033c"

  isSuccessful=$(echo "$result" | grep success)
  areTestsFailed=$(echo "$result" | grep "assertions failed.")
  lineCount=$(echo "$result" | wc -l)

  if [[ $lineCount -eq 2 ]] && [ -n "$isSuccessful" ]; then
      echo "$result
      "
  elif [ -n "$areTestsFailed" ]; then
      printError "$result
      "
      exit 1
  else
      while read -r line; do
          local ln=$(echo "$line" | grep -oh $modulePath":[0-9]*" | sed -e "s=$modulePath:==")
          line=$(echo "$line" | sed -e "s=$modulePath=$1=")

          if [ -z "$ln" ]; then
            echo "$line"
          else
            oln=$(expr "$ln" - 3)
            echo "$line" | sed "s/.go:$ln/.go:$oln/"
          fi

      done <<< "$result"
      exit 1
  fi

}

row () {
  echo -e "    $1"
}

colored () {
  local color="\033[$2m"
  local nc='\033[0m'
  row "${color}$1${nc}"
}

info () {
  colored "$1" "90"
}

title () {
  echo ""
  row "\033[1m$1\033[0m"
  echo ""
}

printError () {
  echo ""
  colored "Error: $1" "31"
  echo ""
  exit 1
}

if [ -n "$command" ]; then
    run "$command"
elif [ "$command" = "help" ]; then
  help $@
else
  help $@
fi
