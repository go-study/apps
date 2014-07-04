<?php
require_once('../parser/pb_parser.php');
$parser = new PBParser();
$parser->parse('./msg.proto');
