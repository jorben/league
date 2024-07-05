import React from 'react'

const Logo = ({collapsed = false, theme = "light"}) => {
    // 定义不同主题的样式
  const themeStyles = {
    light: {
      color: 'rgba(0, 0, 0, 1)',
    //   backgroundColor: 'rgba(255, 255, 255, 0.2)',
    },
    black: {
      color: 'rgba(255, 255, 255, 1)',
    //   backgroundColor: 'rgba(0, 0, 0, 0.2)',
    },
  };
  const logoTextColor = ['#6253E1', '#04BEFE'];
  const currentStyle = themeStyles[theme];
  const pathStyle = {
    stroke: "url(#gradient)",
    strokeWidth: '1.6',
    strokeLinecap: "round",
    strokeLinejoin: "round",
    }

  return (
    <div style={{
        height: '48px',
        lineHeight: '48px',
        ...currentStyle,
        fontSize:'1.5em', 
        margin: '16px', 
        borderRadius: '6px', 
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        }} >
        <svg xmlns="http://www.w3.org/2000/svg" width="36" height="36" viewBox="0 0 24 24" fill="none" >
            <defs>
                <linearGradient id="gradient" x1="0%" y1="0%" x2="100%" y2="0%">
                <stop offset="0%" stopColor="rgb(98, 83, 225)" />
                <stop offset="100%" stopColor="rgb(4, 190, 254)" />
                </linearGradient>
            </defs>
            <path d="M2 2a26.6 26.6 0 0 1 10 20c.9-6.82 1.5-9.5 4-14" style={pathStyle}></path>
            <path d="M16 8c4 0 6-2 6-6-4 0-6 2-6 6" style={pathStyle}></path>
            <path d="M17.41 3.6a10 10 0 1 0 3 3" style={pathStyle}></path>
        </svg>
        <strong style={{
            marginLeft:'8px', 
            backgroundImage:`linear-gradient(135deg, ${logoTextColor.join(', ')})`, 
            backgroundClip:'text', 
            WebkitBackgroundClip: 'text',
            WebkitTextFillColor: 'transparent',
            display: collapsed ? 'none' : ''
            }}>LEAGUE</strong>
      </div>
  )
}

export default Logo
