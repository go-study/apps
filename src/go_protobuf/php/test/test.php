<?php
$data = file_get_contents("php://input");
file_put_contents("/tmp/x.log",$data);
require_once("../message/pb_message.php");
require_once("./pb_proto_msg.php");
//$user = new User();
//$user->set_uid(1);
//$user->set_uname("zhangsan");
//$img1 = new ImgUrl();
//$img1->set_url("root/img/1.jpg");
//$img2 = new ImgUrl();
//$img2->set_url("root/img/2.jpg");
//$img3 = new ImgUrl();
//$img3->set_url("root/img/3.jpg");
//$user->set_imgurl(0,$img1);
//$user->set_imgurl(1,$img2);
//$user->set_imgurl(2,$img3);
//$string = $user->SerializeToString();
//echo "xuliehua : ".$string;       //序列化
$u = new User();
$u->parseFromString($data);
echo "fanxuliehua : ".$u->uname()." uid : ".$u->uid()."  imgurl : ".$u->imgurl(0)->url();       //反序列化
