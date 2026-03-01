# RustDesk API Server Pro锛堝吋瀹瑰寮虹増锛?

[绠€浣撲腑鏂囷紙榛樿锛塢(./README.md) | [English](./README_EN.md)

鍩轰簬寮€婧?[RustDesk](https://github.com/rustdesk/rustdesk) 瀹㈡埛绔皟鐢ㄨ涓虹殑绗笁鏂?API 鏈嶅姟绔疄鐜帮紝骞舵彁渚?Web 绠＄悊鍚庡彴锛坄soybean-admin`锛夈€?

> <span style="color:#ff4d4f;font-weight:700;">警告：本分支部分兼容性更新由 ChatGPT 生成/辅助完成。请在使用前自行审查代码并充分测试，生产环境务必谨慎。</span>

![Dashboard](./img/1.jpeg "Dashboard")

## 鏂囨。涓績锛堝缓璁紭鍏堟煡闃咃級

- [浣跨敤璇存槑锛堥儴缃层€佸垵濮嬪寲銆佸崌绾с€侀獙璇侊級](./docs/USAGE.md)
- [Docker 瀹夎涓庨厤缃弬鑰冿紙瀹夎鍛戒护銆侀粯璁ゅ弬鏁般€佸嵎銆佺鍙ｏ級](./docs/DOCKER.md)
- [绔彛涓庤闂矾寰勮鏄庯紙HTTP/API/Admin/SMTP锛塢(./docs/PORTS.md)
- [闂鎺掓煡鎵嬪唽锛堝父瑙侀棶棰樹笌鏃ュ織瀹氫綅锛塢(./docs/TROUBLESHOOTING.md)
- [鍙戠増璇存槑妯℃澘锛圙itHub Release 鍙洿鎺ュ鐢級](./RELEASE_NOTES.md)

## Docker 閮ㄧ讲绀轰緥锛堝懡浠わ級

```bash
mkdir -p /opt/rustdesk-api-server-pro/data
cd /opt/rustdesk-api-server-pro

curl -L -o server.yaml https://raw.githubusercontent.com/liyan-lucky/rustdesk-api-server-pro/master/backend/server.yaml

docker run -d \\
  --name rustdesk-api-server-pro \\
  --restart unless-stopped \\
  --network host \\
  -e ADMIN_USER=admin \\
  -e ADMIN_PASS='ChangeMe123!' \\
  -v $(pwd)/server.yaml:/app/server.yaml \\
  -v $(pwd)/data:/app/data \\
  ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest

docker logs -f rustdesk-api-server-pro
```

- 璇︾粏 Docker 瀹夎鍛戒护銆丆ompose 绀轰緥銆侀粯璁ゅ弬鏁般€佺鍙ｄ笌鍗疯鏄庯細`docs/DOCKER.md`
- 榛樿绀轰緥浣跨敤 `host` 缃戠粶锛屽疄闄呯洃鍚鍙ｄ互 `server.yaml` 涓?`httpConfig.port` 涓哄噯锛堝父瑙?`:12345`锛?

## 褰撳墠鐘舵€侊紙鍏煎澧炲己鐗堬級

- 褰撳墠鍒嗘敮瀹氫綅涓衡€滃吋瀹瑰寮虹増鈥濓紝鐩爣鏄敖閲忚创杩戞渶鏂扮増 RustDesk 瀹㈡埛绔殑甯哥敤 API 璋冪敤娴佺▼
- 宸插畬鎴愭柊鐗堝鎴风甯歌鎺ュ彛鍏煎琛ラ綈锛堝湴鍧€绨裤€佽澶囧垪琛ㄣ€佸垎缁勯潰鏉垮熀纭€璇锋眰銆佸璁°€乻ysinfo銆乨evices/cli銆乺ecord 绛夛級
- 鍓嶇绠＄悊鍚庡彴宸茬Щ闄よ禐鍔?鏀舵鐮佺浉鍏崇晫闈笌鏂囨
- 浠撳簱涓婚〉榛樿璇█涓轰腑鏂囷紙鍙€氳繃椤堕儴閾炬帴鍒囨崲鑻辨枃璇存槑锛?
- 椤圭洰浠嶄繚鐣欏悗缁噸鏋勮鍒掞紙瑙?issue #30锛夛紝浣嗗綋鍓嶇増鏈彲浠ヤ綔涓衡€滃彲涓婄嚎鍏煎鐗堚€濅娇鐢?

## 椤圭洰瀹氫綅锛堣缁嗚鏄庯級

鏈」鐩笉鏄畼鏂?Pro 鏈嶅姟绔紝涔熶笉杩芥眰鍦ㄧ煭鏈熷唴瀹屽叏澶嶅埢瀹樻柟鍏ㄩ儴楂樼骇鑳藉姏銆傚綋鍓嶇増鏈殑璁捐鐩爣鏄細

- 鍏煎浼樺厛锛氫紭鍏堜繚璇佹柊鐗堝鎴风涓绘祦绋嬩笉鎶ラ敊銆佷笉 404銆佸叧閿瓧娈佃兘姝ｇ‘璇诲啓
- 杞婚噺浼樺厛锛氭敮鎸佸崟鏈洪儴缃诧紝榛樿 SQLite锛岄€傚悎绉佹湁鍖栬嚜寤烘垨涓皬瑙勬ā鍦烘櫙
- 杩唬浼樺厛锛氶€氳繃鍏煎灞傛寔缁ˉ榻愬畼鏂瑰鎴风 API 鍙樺寲锛岄檷浣庡崌绾ф垚鏈?

閫傜敤鍦烘櫙锛?

- 鑷缓 RustDesk API 鏈嶅姟绔?+ 绠＄悊鍚庡彴
- 鍐呯綉/绉佹湁鐜鐨勮澶囩鐞嗕笌鍩虹瀹¤
- 鐮旂┒ RustDesk 瀹㈡埛绔?API 璋冪敤閫昏緫骞跺仛浜屾寮€鍙?

涓嶉€傜敤鎴栭渶璋ㄦ厧璇勪及鐨勫満鏅細

- 寮轰緷璧栧畼鏂?Pro 鐨勫畬鏁?OIDC 鐧诲綍娴佺▼
- 寮轰緷璧栧畼鏂规彃浠剁鍚嶆湇鍔★紙`plugin-sign`锛夌殑鐢熶骇鏍￠獙閾捐矾
- 寮轰緷璧栧畼鏂瑰畬鏁村垎缁勬潈闄愭ā鍨嬶紙`device-group/accessible`銆乣users/peers?accessible=` 鐨勭粏绮掑害鏉冮檺锛?

## 涓庢柊鐗堝鎴风鍏煎鑼冨洿锛堝綋鍓嶄唬鐮佺姸鎬侊級

宸茶鐩栨垨鍏煎澶勭悊鐨勪富瑕佹帴鍙ｄ笌娴佺▼锛?

- 璐﹀彿鐩稿叧锛歚/api/login`銆乣/api/logout`銆乣/api/currentUser`銆乣/api/login-options`
- 鍦板潃绨匡紙鏂版棫鎺ュ彛骞跺瓨锛夛細
  - `/api/ab`
  - `/api/ab/settings`
  - `/api/ab/personal`
  - `/api/ab/shared/profiles`
  - `/api/ab/peers`
  - `/api/ab/tags/{guid}`
  - `/api/ab/peer/*`
  - `/api/ab/tag/*`
- 鍒嗙粍/璁惧闈㈡澘鍩虹璇锋眰锛?
  - `/api/device-group/accessible`
  - `/api/users?accessible=...`
  - `/api/peers?accessible=...`
- 鍚屾涓庣姸鎬侊細
  - `/api/heartbeat`
  - `/api/sysinfo`
  - `/api/sysinfo_ver`
- 瀹¤锛?
  - `/api/audit/conn`
  - `/api/audit/file`
  - `/api/audit/alarm`
  - `PUT /api/audit`锛堝娉ㄦ洿鏂板吋瀹癸級
- 瀹㈡埛绔檮鍔犲吋瀹圭鐐癸細
  - `POST /api/devices/cli`锛堟渶灏忓彲鐢ㄥ啓鍏ワ級
  - `POST /api/record`锛堟渶灏忚惤鐩樺崗璁細`new/part/tail/remove`锛?
  - `POST /api/oidc/auth`锛堝吋瀹硅繑鍥烇級
  - `GET /api/oidc/auth-query`锛堝吋瀹硅繑鍥烇級
  - `POST /lic/web/api/plugin-sign`锛堝吋瀹瑰崰浣嶈繑鍥烇級

## 宸茶ˉ榻愮殑鍏抽敭鍏煎鐐癸紙鐩稿鏃х増绗笁鏂瑰疄鐜帮級

- 鍦板潃绨?`note` 瀛楁鏀寔锛堣/鍐?鍚屾锛?
- 鏂扮増鍦板潃绨垮閲忔洿鏂板瓧娈靛吋瀹癸紙濡?`username`銆乣hostname`銆乣platform`銆乣note`锛?
- `display_name` 瀛楁琛ラ綈锛堢敤鎴?鐧诲綍鐩稿叧鍝嶅簲锛?
- `device_group_name` 瀛楁琛ラ綈锛堣澶囧垪琛ㄥ吋瀹瑰瓧娈碉級
- `devices/cli` 浠?no-op 鍗囩骇涓烘渶灏忓彲鐢ㄥ啓鍏ワ紙璁惧涓庡湴鍧€绨块儴鍒嗗瓧娈靛洖鍐欙級
- `record` 浠?no-op 鍗囩骇涓烘渶灏忚惤鐩樺崗璁疄鐜?
- `sysinfo_ver` 鏀逛负绋冲畾杩斿洖鍊硷紝鍑忓皯瀹㈡埛绔噸澶嶄笂浼?`sysinfo`
- `PUT /api/audit` 澶囨敞鏇存柊鍏煎

## 闈炲畬鏁村畼鏂?Pro 瀹炵幇锛堝彂甯冨墠璇风煡鎮夛級

浠ヤ笅鑳藉姏鐩墠鏄€滃吋瀹瑰疄鐜扳€濇垨鈥滃崰浣嶈繑鍥炩€濓紝鐢ㄤ簬閬垮厤瀹㈡埛绔姤閿欙紝浣嗕笉绛夊悓瀹樻柟 Pro锛?

- OIDC锛坄/api/oidc/*`锛夛細杩斿洖鍏煎缁撴瀯锛屾湭瀹炵幇瀹屾暣鐧诲綍娴佺▼
- 鎻掍欢绛惧悕锛坄/lic/web/api/plugin-sign`锛夛細鍏煎鍗犱綅锛屼笉鏄畼鏂圭鍚嶆湇鍔?
- 鍒嗙粍鏉冮檺妯″瀷锛歚device-group/accessible`銆乣users/peers?accessible=` 涓哄吋瀹规ā鍨嬶紝闈炲畼鏂瑰畬鏁存潈闄愰€昏緫

鍙戝竷鍒ゆ柇寤鸿锛?

- 濡傛灉鐩爣鏄€滄渶鏂扮増瀹㈡埛绔富娴佺▼鍙敤鈥濓細褰撳墠鐗堟湰鍙彂甯?
- 濡傛灉鐩爣鏄€滃畬鍏ㄦ浛浠ｅ畼鏂?Pro 楂樼骇鑳藉姏鈥濓細寤鸿缁х画琛ラ綈 OIDC銆佹彃浠剁鍚嶃€佸畬鏁存潈闄愭ā鍨嬪悗鍐嶅彂甯?

## 鎶€鏈爤

- 鍚庣锛欸o锛圛ris锛?
- 鍓嶇锛歏ue 3 + Vite + Naive UI锛坄soybean-admin`锛?
- 鏁版嵁搴擄細SQLite锛堥粯璁わ級/ MySQL锛堝彲閫夛級

## 椤圭洰缁撴瀯

- `backend/`锛欸o 鍚庣 API 鏈嶅姟
- `soybean-admin/`锛氱鐞嗗悗鍙板墠绔?
- `docker/`锛氬鍣ㄧ浉鍏抽厤缃?
- `docs/`锛氫娇鐢ㄨ鏄庛€佺鍙ｈ鏄庛€侀棶棰樻帓鏌ユ墜鍐?
- `img/`锛歊EADME 鍥剧墖璧勬簮

## 鍙戝竷鍓嶆渶灏忔鏌ワ紙寤鸿锛?

1. 鎵ц鏁版嵁搴撶粨鏋勫悓姝ワ細`rustdesk-api-server-pro.exe sync`
2. 閲嶅惎鏈嶅姟锛岀‘璁?`server.yaml` 涓鍙ｄ笌闈欐€佺洰褰曢厤缃纭?
3. 鐢ㄦ渶鏂扮増瀹㈡埛绔仛涓€杞啋鐑熸祴璇曪紙鐧诲綍銆佸湴鍧€绨裤€佽澶囧垪琛ㄣ€佸垎缁勯潰鏉裤€佸璁★級
4. 妫€鏌ユ棩蹇椾腑鏄惁瀛樺湪鎸佺画鎶ラ敊锛圤IDC/record/鏉冮檺/鏁版嵁搴撳瓧娈电己澶辩瓑锛?

## 璇存槑

- GitHub 椤甸潰鎸夐挳/鏍囩锛堝 `README`銆乣Commits`锛夌殑鏄剧ず璇█鐢?GitHub 涓庢祻瑙堝櫒璇█鍐冲畾锛屼粨搴撲唬鐮佹棤娉曠洿鎺ヤ慨鏀?
- 浠撳簱棣栭〉浠呭睍绀烘€昏淇℃伅锛涜缁嗛儴缃蹭笌鎺掗殰璇锋煡鐪嬩笂鏂光€滄枃妗ｄ腑蹇冣€?
