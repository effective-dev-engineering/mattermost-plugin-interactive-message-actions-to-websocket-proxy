import React, {useCallback, useState, type FC} from 'react';
import styled from 'styled-components';

const MessageContainer = styled.div`
	display: flex;
	align-items: center;
	flex-direction: row;
	gap: 5px;
	padding: 10px 12px;
	background: white;
	border-radius: 4px;
	border: 1px solid rgba(63, 67, 80, 0.08);
`;

const ConfigContainer = styled.div`
	display: flex;
	flex-direction: column;
	gap: 20px;
`;

interface ProxyRule {
	ActionRegExp: string;
	BotUserId:    string;
}

interface Props {
    id: string;
    label: string;
    helpText: JSX.Element | null;
    value: ProxyRule[];
    disabled: boolean;
    config?: Record<string, unknown>;
    license?: Record<string, unknown>;
    setByEnv: boolean;
    onChange: (id: string, value: ProxyRule[], confirm?: boolean, doSubmit?: boolean, warning?: boolean) => void;
    saveAction: () => Promise<unknown>;
    registerSaveAction: (saveAction: () => Promise<{} | {error: {message: string}}>) => void;
    unRegisterSaveAction: (saveAction: () => Promise<unknown>) => void;
    setSaveNeeded: () => void;
    cancelSubmit: () => void;
    showConfirm: boolean;
}

const Config: FC<Props> = ({ id, value, onChange, setSaveNeeded }) => {
    const [proxyRules, setProxyRules] = useState<ProxyRule[]>(value || []);

    const handleChange = useCallback(() => {
        onChange(id, proxyRules);
        setSaveNeeded();
    }, [onChange, setSaveNeeded, proxyRules]);

    return (
        <ConfigContainer>
            <MessageContainer>
                <span>{'1233123123123'}</span>
            </MessageContainer>
        </ConfigContainer>
    );
};

export default Config;
