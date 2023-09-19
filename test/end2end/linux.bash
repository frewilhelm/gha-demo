testCmd () {
	result="$(./gha-demo -zahl1 "$1" -zahl2 "$2" -operator "${operator}" 2>&1)"

	if [[ "${result}" == "$3" ]]; then
		echo "Test passed (${3})"
		return 0
	elif [[ "${result}" =~ ${3} ]]; then
		echo "Test passed (${3})"
		return 0
	else
		echo "Expected '$3', but got '${result}'"
		return 1
	fi
}

rc=0

# Regex for Error-Logging
prefixPattern='[0-9]{4}/[0-9]{2}/[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}'

# Sum
operator='+'

testCmd 2 4 6 || ((rc+=1))
testCmd 99 4 103 || ((rc+=1))

# Error with string in input
testCmd 'a' 4 "^${prefixPattern} gha-demo failed: strconv.Atoi: parsing \"a\": invalid syntax$" || ((rc+=1))
# Error with float in input
testCmd 2.3 4 "^${prefixPattern} gha-demo failed: strconv.Atoi: parsing \"2.3\": invalid syntax$" || ((rc+=1))

# Sub
operator='-'

testCmd 2 4 -2 || ((rc+=1))
testCmd 99 4 95 || ((rc+=1))

# Error with string in input
testCmd 'a' 4 "^${prefixPattern} gha-demo failed: strconv.Atoi: parsing \"a\": invalid syntax$" || ((rc+=1))
# Error with float in input
testCmd 2.3 4 "^${prefixPattern} gha-demo failed: strconv.Atoi: parsing \"2.3\": invalid syntax$" || ((rc+=1))

if [[ $rc -eq 0 ]]; then
	echo 'All tests passed'
	exit 0
else
	echo 'Test(s) failed'

 # Sub
operator='x'

testCmd 2 4 8 || ((rc+=1))
testCmd 99 4 396 || ((rc+=1))

# Error with string in input
testCmd 'a' 4 "^${prefixPattern} gha-demo failed: strconv.Atoi: parsing \"a\": invalid syntax$" || ((rc+=1))
# Error with float in input
testCmd 2.3 4 "^${prefixPattern} gha-demo failed: strconv.Atoi: parsing \"2.3\": invalid syntax$" || ((rc+=1))

if [[ $rc -eq 0 ]]; then
	echo 'All tests passed'
	exit 0
else
	echo 'Test(s) failed'
	exit 1
fi
