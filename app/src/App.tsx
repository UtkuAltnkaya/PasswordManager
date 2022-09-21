import { lazy, Suspense } from 'react';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';
import { useRoutes } from 'react-router-dom';

import Private from 'components/tools/Private';
import Public from 'components/tools/Public';
import Loader from 'components/Loader/Loader';

const Register = lazy(() => import('components/Register/Register'));
const Login = lazy(() => import('components/Login/Login'));
const Main = lazy(() => import('components/Main/Main'));

function App() {
  const element = useRoutes([
    {
      path: '/',
      element: (
        <Suspense fallback={<Loader />}>
          <Private>{(data) => <Main data={data} />}</Private>
        </Suspense>
      ),
    },
    {
      path: '/register',
      element: (
        <Public>
          <Suspense fallback={<Loader />}>
            <Register />
          </Suspense>
        </Public>
      ),
    },
    {
      path: '/login',
      element: (
        <Public>
          <Suspense fallback={<Loader />}>
            <Login />
          </Suspense>
        </Public>
      ),
    },
  ]);

  return (
    <>
      {element}
      <ReactQueryDevtools initialIsOpen={false} />
    </>
  );
}

export default App;
