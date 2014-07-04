<?php
class ImgUrl extends PBMessage
{
  var $wired_type = PBMessage::WIRED_LENGTH_DELIMITED;
  public function __construct($reader=null)
  {
    parent::__construct($reader);
    $this->fields["1"] = "PBString";
    $this->values["1"] = "";
  }
  function url()
  {
    return $this->_get_value("1");
  }
  function set_url($value)
  {
    return $this->_set_value("1", $value);
  }
}
class User extends PBMessage
{
  var $wired_type = PBMessage::WIRED_LENGTH_DELIMITED;
  public function __construct($reader=null)
  {
    parent::__construct($reader);
    $this->fields["1"] = "PBInt";
    $this->values["1"] = "";
    $this->fields["2"] = "PBString";
    $this->values["2"] = "";
    $this->fields["3"] = "ImgUrl";
    $this->values["3"] = array();
  }
  function uid()
  {
    return $this->_get_value("1");
  }
  function set_uid($value)
  {
    return $this->_set_value("1", $value);
  }
  function uname()
  {
    return $this->_get_value("2");
  }
  function set_uname($value)
  {
    return $this->_set_value("2", $value);
  }
  function imgurl($offset)
  {
    return $this->_get_arr_value("3", $offset);
  }
  function add_imgurl()
  {
    return $this->_add_arr_value("3");
  }
  function set_imgurl($index, $value)
  {
    $this->_set_arr_value("3", $index, $value);
  }
  function remove_last_imgurl()
  {
    $this->_remove_last_arr_value("3");
  }
  function imgurl_size()
  {
    return $this->_get_arr_size("3");
  }
}
?>