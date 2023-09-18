function RunTest($input1, $input2, $expect) {
	$result = .\gha-demo.exe -zahl1 $input1 -zahl2 $input2 -operator $operator 2>&1 | Out-String

	$result_trimmed = $result.Trim()

	if ($result_trimmed -eq $expect) {
		Write-Host "Test passed ($expect)"
  		return $true
	} elseif ($result_trimmed.EndsWith($expect)) {
  		Write-Host "Test passed ($expect)"
      		return $true
    	} else {
		Write-Host "Expected '$expect', but got '$result_trimmed'"
    		return $false
	}
}

$rc = 0

# Sum
$operator = '+'

(RunTest 2 4 6) -or $false ? "Test failed" : $rc++
(RunTest 99 4 103) -or $false ? "Test failed" : $rc++

# Error with string in input
(RunTest 'a' 4 'gha-demo failed: strconv.Atoi: parsing "a": invalid syntax') -or $false ? "Test failed" : $rc++
# Error with float in input
(RunTest 2.3 4 'gha-demo failed: strconv.Atoi: parsing "2.3": invalid syntax') -or $false ? "Test failed" : $rc++

if ($rc -eq 0) {
	Write-Host "All tests passed"
 	exit(0)
} else {
	Write-Host "Test(s) failed"
 	exit(1)
}
