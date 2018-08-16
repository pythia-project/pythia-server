#!/usr/bin/env php7.0
<?php
//Create and move to directory
if(!file_exists('/tmp')) {
	mkdir('/tmp', 0666, true);
}

chdir('/tmp');

//Read input and create file
$f = fopen('php://stdin', 'r');
$filename = 'script.php';
$file = fopen($filename, 'a');
chmod($filename, 0777);

while($line=fgets($f)) {
	fwrite($file, $line);
}
fclose($file);

//Send result <missing returncode>
$desc = array(STDIN, array('pipe', "w"), array('pipe','w'));
$process = proc_open('php7.0 -f ' .$filename, $desc, $pipes, '/tmp');
if(is_resource($process)) {
	$output = stream_get_contents($pipes[1]);
	$err = stream_get_contents($pipes[2]);
	$return_val = proc_close($process);
	echo ('{"stdin": "' . $output . '", "stderr": "' .$err . '", "returncode": "' .$return_val. '"}');
}


?>
