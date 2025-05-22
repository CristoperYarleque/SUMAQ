import { redirect } from '@sveltejs/kit';

const publicRoutes = ['/login', '/register'];

export async function handle({ event, resolve }) {
  const token = event.cookies.get('token') || null;
  const role = event.cookies.get('role') || null;
  const path = event.url.pathname;

  // Rutas públicas
  const publicRoutes = ['/login', '/register'];

  // Si no hay token y la ruta no es pública → /login
  if (!token && !publicRoutes.includes(path)) {
    throw redirect(303, '/login');
  }

  // Si hay token y la ruta es pública → dashboard según rol
  if (token && publicRoutes.includes(path)) {
    const home = role === 'entrepreneur' ? '/emprendedor' : '/cliente';
    throw redirect(303, home);
  }

  // Acceso por rol: impedir que un cliente entre a rutas de emprendedor
  if (token && path.startsWith('/emprendedor') && role !== 'entrepreneur') {
    throw redirect(303, '/cliente');
  }

  // Impedir que un emprendedor entre a rutas de cliente
  if (token && path.startsWith('/cliente') && role !== 'client') {
    throw redirect(303, '/emprendedor');
  }

  return resolve(event);
}