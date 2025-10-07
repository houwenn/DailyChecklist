// 主题切换功能
document.addEventListener('DOMContentLoaded', function() {
    const themeToggle = document.getElementById('theme-toggle');
    const themeIcon = document.querySelector('.theme-icon');
    const themeText = document.querySelector('.theme-text');
    
    const isLightTheme = localStorage.getItem('theme') === 'light';
    if (isLightTheme) {
        document.body.classList.add('light-theme');
        themeIcon.textContent = '🌙';
        themeText.textContent = '夜间模式';
    }
    
    themeToggle.addEventListener('click', () => {
        const isLight = document.body.classList.toggle('light-theme');
        
        if (isLight) {
            themeIcon.textContent = '🌙';
            themeText.textContent = '夜间模式';
            localStorage.setItem('theme', 'light');
        } else {
            themeIcon.textContent = '🌞';
            themeText.textContent = '切换主题';
            localStorage.setItem('theme', 'dark');
        }
    });
});

// 图表通用配置
function getCommonChartOptions(titleText) {
    return {
        title: {text: titleText, left: 'center', top: '10px', textStyle: {color: '#4a5568', fontSize: 16, fontWeight: 'normal'}},
        tooltip: {trigger: 'item', backgroundColor: 'rgba(255, 255, 255, 0.95)', borderColor: '#e2e8f0', textStyle: {color: '#2d3748'}},
        legend: {orient: 'horizontal', bottom: '5px', textStyle: {color: '#718096', fontSize: 12}, itemGap: 20}
    };
}