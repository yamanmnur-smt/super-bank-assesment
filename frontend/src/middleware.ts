import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'

function isTokenExpired(token: string): boolean {
  try {
    const payload = JSON.parse(atob(token.split('.')[1]))
    const exp = payload.exp
    const now = Math.floor(Date.now() / 1000)
    return exp < now
  } catch (err) {
    return true 
  }
}

export function middleware(request: NextRequest) {
  const token = request.cookies.get('bearer_token')?.value
  if(!token) {
    
  }

  if ( request.nextUrl.pathname === '/' || request.nextUrl.pathname === '') {
    if (token && !isTokenExpired(token)) {
      const dashboardUrl = new URL('/dashboard', request.url)
      return NextResponse.redirect(dashboardUrl)
    } else {
      const loginUrl = new URL('/login', request.url)
      return NextResponse.redirect(loginUrl)
    }
    
  } else {
    if(token) {
      const isexpired = isTokenExpired(token)
      if(isexpired) {
        const loginUrl = new URL('/login', request.url)
        return NextResponse.redirect(loginUrl)
      } else  {
        if(request.nextUrl.pathname === '/login') {
          const dashboardUrl = new URL('/dashboard', request.url)
          return NextResponse.redirect(dashboardUrl)
        }
      }
    } else {
      const loginUrl = new URL('/login', request.url)
      return NextResponse.redirect(loginUrl)
    }
  }
  
  return NextResponse.next()
}

export const config = {
  matcher: [
    '/dashboard/:path*', 
    '/customers/:path*', 
    '/', 
  ]
}