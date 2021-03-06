/**
 * 演示程序当前的 “注册/登录” 等操作，是基于 “本地存储” 完成的
 * 当您要参考这个演示程序进行相关 app 的开发时，
 * 请注意将相关方法调整成 “基于服务端Service” 的实现。
 **/
(function($, owner) {
	/**
	 * 用户登录
	 **/
	owner.login = function(character, callback) {
		callback = callback || $.noop;
		localStorage.setItem('$users',JSON.stringify(character));
		
		return owner.createState(character,callback);
	};
	
	owner.logout = function() {
		owner.setUser('');	
		owner.setState('');
	};

	owner.createState = function(authRegInfo, callback) {
		var state = owner.getState();
		state.uname = authRegInfo.uname;
		state.token = authRegInfo.token;
		state.nick = authRegInfo.nick;
		state.platform = authRegInfo.platform;
		owner.setState(state);
		return callback();
	};

	/**
	 * 新用户注册
	 **/
	owner.reg = function(regInfo, callback) {
		callback = callback || $.noop;
		regInfo = regInfo || {};
		regInfo.uname = regInfo.uname || '';
		regInfo.password = regInfo.password || '';
		
		var checkUname = /[0-9]/;
		var checkPw = /[a-zA-Z0-9]/
		
		if (!checkEmail(regInfo.uname)) {
			if (regInfo.uname.length != 11 || !checkUname.test(regInfo.uname)) {
				return callback('请输入有效的电话或邮箱')
			}	
		}

		if (regInfo.password.length < 8 || !checkPw.test(regInfo.password)) {
			return callback('密码最短需要 8 个以上的数字或字母');
		}
		if (regInfo.nickname.length < 1) {
			return callback('昵称不能为空');
		}

		return callback();
	};
	
	/**
	 * 获取当前用户信息
	 **/
	owner.getUser = function(){
		var userInfoText = localStorage.getItem('$users') || "{}";
		return JSON.parse(userInfoText);
	};
		/**
	 * 设置当前用户信息
	 **/
	owner.setUser = function(user) {
		userinfo = user || {};
		localStorage.setItem('$users', JSON.stringify(userinfo));
	};
	/**
	 * 获取当前状态
	 **/
	owner.getState = function() {
		var stateText = localStorage.getItem('$state') || "{}";
		return JSON.parse(stateText);
	};

	/**
	 * 设置当前状态
	 **/
	owner.setState = function(state) {
		state = state || {};
		localStorage.setItem('$state', JSON.stringify(state));
	};

	var checkEmail = function(email) {
		email = email || '';
		return (email.length > 3 && email.indexOf('@') > -1);
	};

	/**
	 * 找回密码
	 **/
	owner.forgetPassword = function(email, callback) {
		callback = callback || $.noop;
		if (!checkEmail(email)) {
			return callback('邮箱地址不合法');
		}
		return callback(null, '新的随机密码已经发送到您的邮箱，请查收邮件。');
	};

	/**
	 * 获取应用本地配置
	 **/
	owner.setSettings = function(settings) {
		settings = settings || {};
		localStorage.setItem('$settings', JSON.stringify(settings));
	}

	/**
	 * 设置应用本地配置
	 **/
	owner.getSettings = function() {
			var settingsText = localStorage.getItem('$settings') || "{}";
			return JSON.parse(settingsText);
		}
		/**
		 * 获取本地是否安装客户端
		 **/
	owner.isInstalled = function(id) {
		if (id === 'qihoo' && mui.os.plus) {
			return true;
		}
		if (mui.os.android) {
			var main = plus.android.runtimeMainActivity();
			var packageManager = main.getPackageManager();
			var PackageManager = plus.android.importClass(packageManager)
			var packageName = {
				"qq": "com.tencent.mobileqq",
				"weixin": "com.tencent.mm",
				"sinaweibo": "com.sina.weibo"
			}
			try {
				return packageManager.getPackageInfo(packageName[id], PackageManager.GET_ACTIVITIES);
			} catch (e) {}
		} else {
			switch (id) {
				case "qq":
					var TencentOAuth = plus.ios.import("TencentOAuth");
					return TencentOAuth.iphoneQQInstalled();
				case "weixin":
					var WXApi = plus.ios.import("WXApi");
					return WXApi.isWXAppInstalled()
				case "sinaweibo":
					var SinaAPI = plus.ios.import("WeiboSDK");
					return SinaAPI.isWeiboAppInstalled()
				default:
					break;
			}
		}
	}
}(mui, window.app = {}));