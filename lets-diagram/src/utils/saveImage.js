let save_image_timer = "";

function get_save_image_task(self) {
  return function() {
    if (self.last_push_cover_version < self.version) {
      let image_base64 = window.topology.toImage();
      let now_version = self.version
      let t = self;
      self
        .axios({
          method: "post",
          url: "canvas/cover/update/",
          data: {
            canvas_id: parseInt(t.$route.params.id),
            data: image_base64,
          },
        })
        .then(function(response) {
          if (response.data["header"]["code"] != 200) {
            t.$message.error("封面更新失败");
          } else {
            t.last_push_cover_version = now_version
          }
        });
    }
  };
}

/**
 * 开启一个定时任务，将缩略图定时 push 给服务端
 * @param {*} self
 */
export function start_save_image_task(self) {
  save_image_timer = setInterval(get_save_image_task(self), 1000 * 60);
}

/**
 * 销毁定时任务
 */
export function destory_save_image_timer() {
  clearInterval(save_image_timer);
}

//压缩方法
export function dealImage(base64, w, callback) {
  var newImage = new Image();
  var quality = 0.6; //压缩系数0-1之间
  newImage.src = base64;
  newImage.setAttribute("crossOrigin", "Anonymous"); //url为外域时需要
  var imgWidth, imgHeight;
  newImage.onload = function() {
    imgWidth = this.width;
    imgHeight = this.height;
    var canvas = document.createElement("canvas");
    var ctx = canvas.getContext("2d");
    if (Math.max(imgWidth, imgHeight) > w) {
      if (imgWidth > imgHeight) {
        canvas.width = w;
        canvas.height = (w * imgHeight) / imgWidth;
      } else {
        canvas.height = w;
        canvas.width = (w * imgWidth) / imgHeight;
      }
    } else {
      canvas.width = imgWidth;
      canvas.height = imgHeight;
      quality = 0.6;
    }
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    ctx.drawImage(this, 0, 0, canvas.width, canvas.height);
    var base64 = canvas.toDataURL("image/jpeg", quality); //压缩语句
    callback(base64); //必须通过回调函数返回，否则无法及时拿到该值
  };
}
