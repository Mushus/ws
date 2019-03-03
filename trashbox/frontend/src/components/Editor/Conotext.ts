import { createContext } from 'react';

const Context = createContext({ selected: [] } as { selected: string[] });

export default Context;
