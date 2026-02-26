import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: { ...enUs.system, title: 'Rustdesk Api Server' },
  common: {
    ...enUs.common,
    refresh: '새로고침',
    noData: '데이터 없음',
    error: '오류',
    warning: '경고'
  },
  route: {
    ...enUs.route,
    home: '홈',
    audit: '감사',
    user: '사용자 관리',
    user_list: '사용자 목록',
    user_sessions: '세션',
    system: '시스템 관리',
    devices: '장치'
  },
  page: {
    ...enUs.page,
    home: {
      ...enUs.page.home,
      greeting: '좋은 아침입니다, {userName}님!',
      userCount: '사용자 수',
      deviceCount: '장치 수',
      onlineCount: '온라인 수',
      visitsCount: '방문 수',
      operatingSystem: '운영체제',
      oneWeek: '최근 1주',
      changeLogs: '업데이트 로그',
      serverConfig: {
        ...enUs.page.home.serverConfig,
        title: '클라이언트 연결 설정',
        tip: '아래 설정을 RustDesk 클라이언트에 붙여넣으세요. KEY가 비어 있으면 `RUSTDESK_KEY` 환경변수를 설정하세요.',
        refresh: '새로고침',
        copyAll: '전체 복사',
        copyTemplate: 'RustDesk 템플릿 복사',
        connectivity: {
          ...enUs.page.home.serverConfig.connectivity,
          check: '연결 확인',
          checkOne: '확인',
          clear: '결과 지우기'
        }
      }
    }
  },
  icon: { ...enUs.icon, lang: '언어 전환', reload: '페이지 새로고침', fullscreen: '전체화면' }
};

export default local;
