<?php
require 'vendor/autoload.php';
use PhpParser\Error;
use PhpParser\ParserFactory;

$code   = file_get_contents('/Users/wangkun/localhost/test/test.php');
$parser = (new ParserFactory)->create(ParserFactory::PREFER_PHP7);
try {
    $ast = $parser->parse($code);
} catch (Error $error) {
    echo "Parse error: {$error->getMessage()}\n";
    return;
}
echo json_encode($ast, JSON_PRETTY_PRINT);
