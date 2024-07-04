import React from 'react'
import Icon from './Icon'

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
        <Icon.Vegan size={36} strokeWidth={1.5} style={{lineHeight: '48px'}} />
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
