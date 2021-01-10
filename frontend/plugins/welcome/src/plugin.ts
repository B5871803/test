import { createPlugin } from '@backstage/core';
import Operativerecord from './components/Operativerecord';
 
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/Operativerecord', Operativerecord);
  },
});
