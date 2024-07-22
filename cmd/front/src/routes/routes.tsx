import { createBrowserRouter, RouterProvider } from 'react-router-dom';

import ErrorPage from '@/errors/error-page';
import HomePage from '@/routes/home-page';
import CollectionPage from '@/routes/collection-page';
import CollectionRun from './collection-run-page';
import CollectionAdd from '@/routes/collection-add-page';
import CollectionDetail from './collection-detail-page';

const Routes = () => {

	const anonymousRoutes = [
		{
			path: '/',
			errorElement: <ErrorPage />,
			children: [
				{
					path: '/',
					element: <HomePage />,
				},
				{
					path: 'collection',
          children:[
            {
              path: '',
              element: <CollectionPage />,
            },
            {
              path: ':id',
              element: <CollectionDetail />,
            },
            {
              path: 'add',
              element: <CollectionAdd />,
            },
            {
              path: 'run/:id',
              element: <CollectionRun />,
            },
          ]
				},
			],
		},
	];

	

	const router = createBrowserRouter([...anonymousRoutes]);

	return <RouterProvider router={router} />;
};

export default Routes;
